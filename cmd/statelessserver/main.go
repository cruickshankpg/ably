package main

import (
	"ably/internal/stateless"
	"context"
	"flag"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	host := flag.String("host", "localhost", "address to listen to")
	port := flag.Int("port", 9001, "port to listen on")
	flag.Parse()

	address := fmt.Sprintf("%s:%d", *host, *port)

	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime: "@timestamp",
		},
	})
	log.Info("starting stateless server listening on " + address)

	runCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, os.Interrupt, os.Kill, syscall.SIGTERM)
		sig := <-sigCh
		log.WithField("signal", sig.String()).Info("received os signal to shutdown")
		cancel()
	}()

	if err := stateless.New(log, time.Now().Unix()).Run(runCtx, address); err != nil {
		log.WithError(err).Error("server failed")
		return
	}

	log.Info("shutting down")
}
