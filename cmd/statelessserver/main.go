package main

import (
	"ably/internal/stateless"
	"context"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime: "@timestamp",
		},
	})
	log.Info("starting stateless server")

	runCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, os.Interrupt, os.Kill, syscall.SIGTERM)
		sig := <-sigCh
		log.WithField("signal", sig.String()).Info("received os signal to shutdown")
		cancel()
	}()

	if err := stateless.New(log, time.Now().Unix()).Run(runCtx, "localhost:9001"); err != nil {
		log.WithError(err).Error("server failed")
		return
	}

	log.Info("shutting down")
}
