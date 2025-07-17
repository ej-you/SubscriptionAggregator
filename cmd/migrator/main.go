// Migrator binary is a migration manager for server DB.
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	cli "github.com/urfave/cli/v3"

	"SubscriptionAggregator/cmd/migrator/commands"
	"SubscriptionAggregator/config"
	"SubscriptionAggregator/internal/pkg/migrate"
)

func main() {
	if err := startMigrator(); err != nil {
		logrus.Fatal(err)
	}
}

func startMigrator() error {
	// load config
	cfg, err := config.New()
	if err != nil {
		return err
	}
	// create migrate manager
	migrateManager, err := migrate.NewCockroachMigrate(
		cfg.DB.MigrationsURL, cfg.DB.ConnURL)
	if err != nil {
		return err
	}
	// defer migrate manager close
	defer migrateManager.Close()

	// create migrator cmd
	cmd := &cli.Command{
		Name:  "migrator",
		Usage: "Migration manager for application DB",
		Commands: []*cli.Command{
			commands.NewStatus(migrateManager),
			commands.NewDown(migrateManager),
			commands.NewUp(migrateManager),
			commands.NewForce(migrateManager),
		},
	}
	// run migrator cmd
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		return fmt.Errorf("migrator cmd: %w", err)
	}
	return nil
}
