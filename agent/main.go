package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/kube-tarian/container-bridge/agent/pkg/application"
)

func main() {
	app := application.New()
	go app.Start()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	<-signals

	app.Close()
}
