package command

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path"

	userpb "github.com/cs3org/go-cs3apis/cs3/identity/user/v1beta1"
	"github.com/cs3org/reva/v2/cmd/revad/runtime"
	"github.com/gofrs/uuid"
	"github.com/oklog/run"
	"github.com/owncloud/ocis/extensions/storage-system/pkg/config"
	"github.com/owncloud/ocis/extensions/storage-system/pkg/config/parser"
	"github.com/owncloud/ocis/extensions/storage/pkg/server/debug"
	"github.com/owncloud/ocis/extensions/storage/pkg/service/external"
	ociscfg "github.com/owncloud/ocis/ocis-pkg/config"
	"github.com/owncloud/ocis/ocis-pkg/log"
	"github.com/owncloud/ocis/ocis-pkg/sync"
	"github.com/owncloud/ocis/ocis-pkg/tracing"
	"github.com/owncloud/ocis/ocis-pkg/version"
	"github.com/thejerf/suture/v4"
	"github.com/urfave/cli/v2"
)

// StorageSystem the entrypoint for the storage-system command.
//
// It provides a ocis-specific storage store system data (shares,account,settings...)
func StorageSystem(cfg *config.Config) *cli.Command {
	return &cli.Command{
		Name:     "storage-system",
		Usage:    "start storage-system service",
		Category: "extensions",
		Before: func(ctx *cli.Context) error {
			err := parser.ParseConfig(cfg)
			if err != nil {
				fmt.Printf("%v", err)
			}
			return err
		},
		Action: func(c *cli.Context) error {
			logCfg := cfg.Logging
			logger := log.NewLogger(
				log.Level(logCfg.Level),
				log.File(logCfg.File),
				log.Pretty(logCfg.Pretty),
				log.Color(logCfg.Color),
			)
			tracing.Configure(cfg.Tracing.Enabled, cfg.Tracing.Type, logger)

			gr := run.Group{}
			ctx, cancel := func() (context.Context, context.CancelFunc) {
				if cfg.Context == nil {
					return context.WithCancel(context.Background())
				}
				return context.WithCancel(cfg.Context)
			}()

			defer cancel()

			pidFile := path.Join(os.TempDir(), "revad-"+c.Command.Name+"-"+uuid.Must(uuid.NewV4()).String()+".pid")
			rcfg := storageSystemFromStruct(c, cfg)

			gr.Add(func() error {
				runtime.RunWithOptions(
					rcfg,
					pidFile,
					runtime.WithLogger(&logger.Logger),
				)
				return nil
			}, func(_ error) {
				logger.Info().
					Str("server", c.Command.Name).
					Msg("Shutting down server")

				cancel()
			})

			debugServer, err := debug.Server(
				debug.Name(c.Command.Name+"-debug"),
				debug.Addr(cfg.Debug.Addr),
				debug.Logger(logger),
				debug.Context(ctx),
				debug.Pprof(cfg.Debug.Pprof),
				debug.Zpages(cfg.Debug.Zpages),
				debug.Token(cfg.Debug.Token),
			)

			if err != nil {
				logger.Info().
					Err(err).
					Str("server", c.Command.Name+"-debug").
					Msg("Failed to initialize server")

				return err
			}

			gr.Add(func() error {
				return debugServer.ListenAndServe()
			}, func(_ error) {
				_ = debugServer.Shutdown(ctx)
				cancel()
			})

			if !cfg.Supervised {
				sync.Trap(&gr, cancel)
			}

			if err := external.RegisterGRPCEndpoint(
				ctx,
				"com.owncloud.storage.system",
				uuid.Must(uuid.NewV4()).String(),
				cfg.GRPC.Addr,
				version.String,
				logger,
			); err != nil {
				logger.Fatal().Err(err).Msg("failed to register the grpc endpoint")
			}

			return gr.Run()
		},
	}
}

// storageSystemFromStruct will adapt an oCIS config struct into a reva mapstructure to start a reva service.
func storageSystemFromStruct(c *cli.Context, cfg *config.Config) map[string]interface{} {
	rcfg := map[string]interface{}{
		"core": map[string]interface{}{
			"tracing_enabled":      cfg.Tracing.Enabled,
			"tracing_endpoint":     cfg.Tracing.Endpoint,
			"tracing_collector":    cfg.Tracing.Collector,
			"tracing_service_name": c.Command.Name,
		},
		"shared": map[string]interface{}{
			"jwt_secret":                cfg.TokenManager.JWTSecret,
			"gatewaysvc":                cfg.Reva.Address,
			"skip_user_groups_in_token": cfg.SkipUserGroupsInToken,
		},
		"grpc": map[string]interface{}{
			"network": cfg.GRPC.Protocol,
			"address": cfg.GRPC.Addr,
			"services": map[string]interface{}{
				"gateway": map[string]interface{}{
					// registries are located on the gateway
					"authregistrysvc":    cfg.GRPC.Addr,
					"storageregistrysvc": cfg.GRPC.Addr,
					// user metadata is located on the users services
					"userprovidersvc":  cfg.GRPC.Addr,
					"groupprovidersvc": cfg.GRPC.Addr,
					"permissionssvc":   cfg.GRPC.Addr,
					// other
					"disable_home_creation_on_login": true, // system manually creates a space
					// system always uses the simple upload, so no transfer secret or datagateway needed
				},
				"userprovider": map[string]interface{}{
					"driver": "memory",
					"drivers": map[string]interface{}{
						"memory": map[string]interface{}{
							"users": map[string]interface{}{
								"serviceuser": map[string]interface{}{
									"id": map[string]interface{}{
										"opaqueId": cfg.SystemUserID,
										"idp":      "internal",
										"type":     userpb.UserType_USER_TYPE_PRIMARY,
									},
									"username":     "serviceuser",
									"display_name": "System User",
								},
							},
						},
					},
				},
				"authregistry": map[string]interface{}{
					"driver": "static",
					"drivers": map[string]interface{}{
						"static": map[string]interface{}{
							"rules": map[string]interface{}{
								"machine": cfg.GRPC.Addr,
							},
						},
					},
				},
				"authprovider": map[string]interface{}{
					"auth_manager": "machine",
					"auth_managers": map[string]interface{}{
						"machine": map[string]interface{}{
							"api_key":      cfg.MachineAuthAPIKey,
							"gateway_addr": cfg.GRPC.Addr,
						},
					},
				},
				"permissions": map[string]interface{}{
					"driver": "demo",
					"drivers": map[string]interface{}{
						"demo": map[string]interface{}{},
					},
				},
				"storageregistry": map[string]interface{}{
					"driver": "static",
					"drivers": map[string]interface{}{
						"static": map[string]interface{}{
							"rules": map[string]interface{}{
								"/": map[string]interface{}{
									"address": cfg.GRPC.Addr,
								},
							},
						},
					},
				},
				"storageprovider": map[string]interface{}{
					"driver":          cfg.Driver,
					"drivers":         config.StorageSystemDrivers(cfg),
					"data_server_url": cfg.DataServerURL,
					"tmp_folder":      cfg.TempFolder,
				},
			},
		},
		"http": map[string]interface{}{
			"network": cfg.HTTP.Protocol,
			"address": cfg.HTTP.Addr,
			// no datagateway needed as the system client directly talks to the dataprovider with the simple protocol
			"services": map[string]interface{}{
				"dataprovider": map[string]interface{}{
					"prefix":      "data",
					"driver":      cfg.Driver,
					"drivers":     config.StorageSystemDrivers(cfg),
					"timeout":     86400,
					"insecure":    cfg.DataProviderInsecure,
					"disable_tus": true,
				},
			},
		},
	}
	return rcfg
}

// SutureService allows for the storage-system command to be embedded and supervised by a suture supervisor tree.
type SystemSutureService struct {
	cfg *config.Config
}

// NewSutureService creates a new storagesystem.SutureService
func NewStorageSystem(cfg *ociscfg.Config) suture.Service {
	cfg.StorageSystem.Commons = cfg.Commons
	return SystemSutureService{
		cfg: cfg.StorageSystem,
	}
}

func (s SystemSutureService) Serve(ctx context.Context) error {
	s.cfg.Context = ctx
	f := &flag.FlagSet{}
	cmdFlags := StorageSystem(s.cfg).Flags
	for k := range cmdFlags {
		if err := cmdFlags[k].Apply(f); err != nil {
			return err
		}
	}
	cliCtx := cli.NewContext(nil, f, nil)
	if StorageSystem(s.cfg).Before != nil {
		if err := StorageSystem(s.cfg).Before(cliCtx); err != nil {
			return err
		}
	}
	if err := StorageSystem(s.cfg).Action(cliCtx); err != nil {
		return err
	}

	return nil
}
