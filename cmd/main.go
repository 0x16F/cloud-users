package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/0x16F/cloud/users/internal/controller/httpsrv"
	"github.com/0x16F/cloud/users/internal/definitions"
	"github.com/0x16F/cloud/users/internal/usecase/config"
	"github.com/0x16F/cloud/users/pkg/logger"
)

func main() {
	container, err := definitions.New()
	if err != nil {
		panic(err)
	}

	cfg, _ := container.Get(definitions.ConfigDef).(*config.Config)
	server, _ := container.Get(definitions.HTTPServerDef).(*httpsrv.Server)
	log, _ := container.Get(definitions.LoggerDef).(logger.Logger)

	go func() {
		if err := server.Start(cfg.WebServer.Port); err != nil {
			log.Fatalf("failed to start server: %v", err)
		}
	}()

	run()

	container.Delete()
}

func run() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	<-signalChan
}
