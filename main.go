// Package main service entrypoint
package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/trevatk/go-pkg/logging"
	"github.com/trevatk/go-template/internal/port"
)

func main() {

	fxApp := fx.New(
		fx.Provide(logging.New),
		fx.Provide(port.NewHTTPServer),
		fx.Provide(fx.Annotate(port.NewRouter, fx.As(new(http.Handler)))),
		fx.Invoke(registerHooks),
	)

	start, cancel := context.WithTimeout(context.TODO(), time.Second*15)
	defer cancel()

	if err := fxApp.Start(start); err != nil {
		log.Fatalf("error starting service %v", err)
	}

	<-fxApp.Done()

	stop, cancel := context.WithTimeout(context.TODO(), time.Second*15)
	defer cancel()

	if err := fxApp.Stop(stop); err != nil {
		log.Fatalf("error stopping service %v", err)
	}
}

func registerHooks(lc fx.Lifecycle, log *zap.Logger, handler http.Handler) error {

	l := log.Named("lifecycle").Sugar()

	p := os.Getenv("HTTP_SERVER_PORT")
	if p == "" {
		return errors.New("$HTTP_SERVER_PORT is unset")
	}

	srv := &http.Server{
		Addr:         ":" + p,
		Handler:      handler,
		ReadTimeout:  time.Second * 15,
		WriteTimeout: time.Second * 15,
		IdleTimeout:  time.Second * 15,
	}

	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {

				l.Infof("start http server http://localhost:%s", p)

				go func() {
					if err := srv.ListenAndServe(); err != nil {
						l.Fatalf("failed to start http server %v", err)
					}
				}()

				return nil
			},
			OnStop: func(ctx context.Context) error {

				var err error

				l.Info("shutdown http server")

				err = srv.Close()
				if err != nil && !errors.Is(err, http.ErrServerClosed) {
					l.Errorf("failed to shutdown http server %v", err)
				}

				// redudant logging
				return err
			},
		},
	)

	return nil
}
