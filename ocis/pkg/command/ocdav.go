package command

import (
	"fmt"

	"github.com/owncloud/ocis/v2/ocis-pkg/config"
	"github.com/owncloud/ocis/v2/ocis-pkg/config/parser"
	"github.com/owncloud/ocis/v2/ocis/pkg/command/helper"
	"github.com/owncloud/ocis/v2/ocis/pkg/register"
	"github.com/owncloud/ocis/v2/services/ocdav/pkg/command"
	"github.com/urfave/cli/v2"
)

// OCDavCommand is the entrypoint for the OCDav command.
func OCDavCommand(cfg *config.Config) *cli.Command {
	return &cli.Command{
		Name:     cfg.OCDav.Service.Name,
		Usage:    helper.SubcommandDescription(cfg.OCDav.Service.Name),
		Category: "services",
		Before: func(c *cli.Context) error {
			if err := parser.ParseConfig(cfg, true); err != nil {
				fmt.Printf("%v", err)
			}
			cfg.OCDav.Commons = cfg.Commons
			return nil
		},
		Subcommands: command.GetCommands(cfg.OCDav),
	}
}

func init() {
	register.AddCommand(OCDavCommand)
}
