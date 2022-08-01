// Package app configures and runs application.
package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"

	"glogin/config"
	v1 "glogin/internal/controller/http/v1"
	"glogin/internal/usecase"
	"glogin/internal/usecase/repo"
	"glogin/pkg/httpserver"
	"glogin/pkg/logger"
	"glogin/pkg/postgres"
)

// Run creates objects via constructors.
func Run(cfg *config.Config, port string) {
	l := logger.New(cfg.Log.Level)

	// Repository
	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	} else {
		l.Info("app - Run - postgres.")
	}
	defer pg.Close()

	// Use case
	loginUserCase := usecase.New(repo.New(pg), cfg.App.TokenExpire, cfg.App.Secret)
	// HTTP Server
	handler := gin.New()
	v1.NewRouter(handler, l, loginUserCase)
	httpServer := httpserver.New(handler, httpserver.Port(port))
	l.Info("app - Run - httpServer: " + port + ".")

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
		// case err = <-rmqServer.Notify():
		// 	l.Error(fmt.Errorf("app - Run - rmqServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}

}
