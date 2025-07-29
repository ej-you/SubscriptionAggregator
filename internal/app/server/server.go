// Package server provides HTTP-server interface.
package server

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"SubscriptionAggregator/config"
	"SubscriptionAggregator/internal/app/errors"
	"SubscriptionAggregator/internal/app/middleware"

	"SubscriptionAggregator/internal/pkg/database"
	"SubscriptionAggregator/internal/pkg/jsonify"
	"SubscriptionAggregator/internal/pkg/logger"
	"SubscriptionAggregator/internal/pkg/validator"
)

var _ Server = (*httpServer)(nil)

// HTTP-server interface.
type Server interface {
	Run()
	WaitForShutdown() error
}

// HTTP-server implementation.
type httpServer struct {
	cfg     *config.Config
	db      *gorm.DB
	valid   validator.Validator
	jsonify jsonify.Jsonify

	fiberApp *fiber.App
	err      chan error // server listen error
}

// New returns new Server instance.
func New(cfg *config.Config) (Server, error) {
	logger.Init(cfg.App.LogLevel, cfg.App.LogFormat)

	gormDB, err := database.New(cfg.DB.ConnString,
		database.WithTranslateError(),
		database.WithIgnoreNotFound(),
		database.WithLogLevel(cfg.App.LogLevel),
		database.WithLogger(logrus.StandardLogger()),
	)
	if err != nil {
		return nil, fmt.Errorf("db: %w", err)
	}

	return &httpServer{
		cfg:     cfg,
		db:      gormDB,
		valid:   validator.New(),
		jsonify: jsonify.New(),
		err:     make(chan error),
	}, nil
}

//	@title			Subscription Aggregator API
//	@version		1.0.0
//	@description	HTTP API для агрегации данных об онлайн-подписках пользователей
//
//	@host			127.0.0.1:8000
//	@basePath		/api/v1
//	@schemes		http
//
//	@accept			json
//	@produce		json
//
// Run starts server.
func (s *httpServer) Run() {
	// app init
	s.fiberApp = fiber.New(fiber.Config{
		AppName:       s.cfg.Server.Name,
		ErrorHandler:  errors.CustomErrorHandler,
		JSONEncoder:   s.jsonify.Marshal,
		JSONDecoder:   s.jsonify.Unmarshal,
		ServerHeader:  "Subscription Aggregator API",
		StrictRouting: false,
	})

	// set up base middlewares
	httpLogger := middleware.Logger(s.cfg.App.LogLevel, s.cfg.App.LogFormat)
	if httpLogger != nil {
		s.fiberApp.Use(httpLogger)
	}
	s.fiberApp.Use(middleware.Recover())
	s.fiberApp.Use(middleware.Swagger())
	// register all endpoints
	s.registerEndpointsV1()

	// start app
	go func() {
		if err := s.fiberApp.Listen(":" + s.cfg.Server.Port); err != nil {
			s.err <- fmt.Errorf("listen: %w", err)
		}
	}()
}

// WaitForShutdown waits for OS signal to gracefully shuts down server.
// This method is blocking.
func (s *httpServer) WaitForShutdown() error {
	// skip if server is not running
	if s.fiberApp == nil {
		return nil
	}

	// handle shutdown process signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	shutdownDone := make(chan struct{})
	// create gracefully shutdown task
	var err error
	go func() {
		defer close(shutdownDone)
		select {
		case err = <-s.err: // server listen error
			return
		case handledSignal := <-quit:
			logrus.Infof("Got %s signal. Shutdown server", handledSignal.String())
			// shutdown app
			s.fiberApp.ShutdownWithTimeout(s.cfg.Server.ShutdownTimeout) // nolint:errcheck // cannot occurs
		}
	}()

	// wait for shutdown
	<-shutdownDone
	logrus.Info("Server shutdown successfully")
	return err
}
