package main

import (
	"ably/internal/stateful"
	"ably/internal/store"
	"context"
	"flag"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	host := flag.String("host", "localhost", "address to listen to")
	port := flag.Int("port", 9000, "port to listen on")
	debug := flag.Bool("debug", false, "close stream after each message")
	flag.Parse()
	address := fmt.Sprintf("%s:%d", *host, *port)
	log := logrus.New()

	log.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime: "@timestamp",
		},
	})
	log.Info("starting stateful server listening on " + address)

	runCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, os.Interrupt, os.Kill, syscall.SIGTERM)
		sig := <-sigCh
		log.WithField("signal", sig.String()).Info("received os signal to shutdown")
		cancel()
	}()

	if err := stateful.New(log, store.New(), *debug).Run(runCtx, address); err != nil {
		log.WithError(err).Error("server failed")
		return
	}

	log.Info("shutting down")
}
