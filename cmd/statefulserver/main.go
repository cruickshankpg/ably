package main

import (
	"ably/internal/stateful"
	"context"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime: "@timestamp",
		},
	})
	log.Info("starting stateful server")

	runCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, os.Interrupt, os.Kill, syscall.SIGTERM)
		sig := <-sigCh
		log.WithField("signal", sig.String()).Info("received os signal to shutdown")
		cancel()
	}()

	if err := stateful.New(log, struct{}{}).Run(runCtx, "localhost:9000"); err != nil {
		log.WithError(err).Error("server failed")
		return
	}

	log.Info("shutting down")
}
