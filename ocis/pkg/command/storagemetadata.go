package command

import (
	"fmt"

	"github.com/owncloud/ocis/extensions/storage-metadata/pkg/command"
	"github.com/owncloud/ocis/ocis-pkg/config"
	"github.com/owncloud/ocis/ocis-pkg/config/parser"
	"github.com/owncloud/ocis/ocis/pkg/register"
	"github.com/urfave/cli/v2"
)

// StorageMetadataCommand is the entrypoint for the storage-metadata command.
func StorageMetadataCommand(cfg *config.Config) *cli.Command {
	return &cli.Command{
		Name:     "storage-metadata",
		Usage:    "start storage and data service for metadata",
		Category: "extensions",
		Before: func(c *cli.Context) error {
			if err := parser.ParseConfig(cfg); err != nil {
				fmt.Printf("%v", err)
				return err
			}
			cfg.StorageMetadata.Commons = cfg.Commons
			return nil
		},
		Action: func(c *cli.Context) error {
			origCmd := command.StorageMetadata(cfg.StorageMetadata)
			return handleOriginalAction(c, origCmd)
		},
	}
}

func init() {
	register.AddCommand(StorageMetadataCommand)
}
