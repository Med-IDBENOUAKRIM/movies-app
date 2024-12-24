package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/med-IDBENOUAKRIM/lets_go/cmd/utils"
)

func (app *Application) serve() error {
	utils.LoadConfig()
	port := os.Getenv("SERVER_ADDRESS")
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	srv := &http.Server{
		Addr:         fmt.Sprintf("localhost:%s", port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	go func() {
		quit := make(chan os.Signal, 1)

		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		s := <-quit

		app.logger.Info("caught signal", "signal", s.String())

		os.Exit(0)
	}()

	logger.Info("starting server", "addr", srv.Addr, "env", app.config.env)
	return srv.ListenAndServe()

}

// ! 259
