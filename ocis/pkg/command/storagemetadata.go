package command

import (
	"fmt"

	"github.com/owncloud/ocis/extensions/storage-system/pkg/command"
	"github.com/owncloud/ocis/ocis-pkg/config"
	"github.com/owncloud/ocis/ocis-pkg/config/parser"
	"github.com/owncloud/ocis/ocis/pkg/register"
	"github.com/urfave/cli/v2"
)

// StorageSystemCommand is the entrypoint for the storage-system command.
func StorageSystemCommand(cfg *config.Config) *cli.Command {
	return &cli.Command{
		Name:     "storage-system",
		Usage:    "start storage and data service for system metadata",
		Category: "extensions",
		Before: func(c *cli.Context) error {
			if err := parser.ParseConfig(cfg); err != nil {
				fmt.Printf("%v", err)
				return err
			}
			cfg.StorageSystem.Commons = cfg.Commons
			return nil
		},
		Action: func(c *cli.Context) error {
			origCmd := command.StorageSystem(cfg.StorageSystem)
			return handleOriginalAction(c, origCmd)
		},
	}
}

func init() {
	register.AddCommand(StorageSystemCommand)
}
