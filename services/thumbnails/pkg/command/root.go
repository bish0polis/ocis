package command

import (
	"context"
	"os"

	"github.com/owncloud/ocis/v2/ocis-pkg/clihelper"
	ociscfg "github.com/owncloud/ocis/v2/ocis-pkg/config"
	"github.com/owncloud/ocis/v2/services/thumbnails/pkg/config"
	"github.com/thejerf/suture/v4"
	"github.com/urfave/cli/v2"
)

// GetCommands provides all commands for this service
func GetCommands(cfg *config.Config) cli.Commands {
	return []*cli.Command{
		// start this service
		Server(cfg),

		// interaction with this service

		// infos about this service
		Health(cfg),
		Version(cfg),
	}
}

// Execute is the entry point for the ocis-thumbnails command.
func Execute(cfg *config.Config) error {
	app := clihelper.DefaultApp(&cli.App{
		Name:     "thumbnails",
		Usage:    "Example usage",
		Commands: GetCommands(cfg),
	})

	return app.Run(os.Args)
}

// SutureService allows for the thumbnails command to be embedded and supervised by a suture supervisor tree.
type SutureService struct {
	cfg *config.Config
}

// NewSutureService creates a new thumbnails.SutureService
func NewSutureService(cfg *ociscfg.Config) suture.Service {
	cfg.Thumbnails.Commons = cfg.Commons
	return SutureService{
		cfg: cfg.Thumbnails,
	}
}

func (s SutureService) Serve(ctx context.Context) error {
	s.cfg.Context = ctx
	if err := Execute(s.cfg); err != nil {
		return err
	}

	return nil
}