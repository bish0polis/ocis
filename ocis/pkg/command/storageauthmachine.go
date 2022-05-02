package command

import (
	"fmt"

	"github.com/owncloud/ocis/extensions/auth-machine/pkg/command"
	"github.com/owncloud/ocis/ocis-pkg/config"
	"github.com/owncloud/ocis/ocis-pkg/config/parser"
	"github.com/owncloud/ocis/ocis/pkg/register"
	"github.com/urfave/cli/v2"
)

// StorageAuthMachineCommand is the entrypoint for the reva-auth-machine command.
func StorageAuthMachineCommand(cfg *config.Config) *cli.Command {
	return &cli.Command{
		Name:     "storage-auth-machine",
		Usage:    "start storage auth-machine service",
		Category: "extensions",
		Before: func(c *cli.Context) error {
			if err := parser.ParseConfig(cfg); err != nil {
				fmt.Printf("%v", err)
				return err
			}
			cfg.AuthMachine.Commons = cfg.Commons
			return nil
		},
		Action: func(c *cli.Context) error {
			origCmd := command.AuthMachine(cfg.AuthMachine)
			return handleOriginalAction(c, origCmd)
		},
	}
}

func init() {
	register.AddCommand(StorageAuthMachineCommand)
}
