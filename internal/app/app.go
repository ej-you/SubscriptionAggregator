// Package app provides function Run to start full application.
package app

import (
	"fmt"

	"SubscriptionAggregator/config"
	"SubscriptionAggregator/internal/app/server"
)

// Run loads app config and starts HTTP-server.
// This function is blocking.
func Run() error {
	// create config
	cfg, err := config.New()
	if err != nil {
		return fmt.Errorf("create config: %w", err)
	}
	// create HTTP-server
	srv, err := server.New(cfg)
	if err != nil {
		return fmt.Errorf("create server: %w", err)
	}
	// run HTTP-server and wait for HTTP-server shutdown
	srv.Run()
	if err := srv.WaitForShutdown(); err != nil {
		return fmt.Errorf("http server: %w", err)
	}
	return nil
}
