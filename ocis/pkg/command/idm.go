package command

import (
	"fmt"

	"github.com/owncloud/ocis/v2/ocis-pkg/config"
	"github.com/owncloud/ocis/v2/ocis-pkg/config/parser"
	"github.com/owncloud/ocis/v2/ocis/pkg/command/helper"
	"github.com/owncloud/ocis/v2/ocis/pkg/register"
	"github.com/owncloud/ocis/v2/services/idm/pkg/command"
	"github.com/urfave/cli/v2"
)

// IDMCommand is the entrypoint for the idm command.
func IDMCommand(cfg *config.Config) *cli.Command {
	return &cli.Command{
		Name:     cfg.IDM.Service.Name,
		Usage:    helper.SubcommandDescription(cfg.IDM.Service.Name),
		Category: "services",
		Before: func(c *cli.Context) error {
			if err := parser.ParseConfig(cfg, true); err != nil {
				fmt.Printf("%v", err)
			}
			cfg.IDM.Commons = cfg.Commons
			return nil
		},
		Subcommands: command.GetCommands(cfg.IDM),
	}
}

func init() {
	register.AddCommand(IDMCommand)
}
