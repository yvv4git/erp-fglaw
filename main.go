package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/yvv4git/erp-fglaw/internal/config"
)

const (
	configFile = "config/main"
)

func main() {
	log.Println("Start server")

	// Init config.
	cfg, err := config.Init(configFile)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(cfg)

	// For graceful shutdown.
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Gracefull exit.
	<-exit
	log.Println("Stopping app...")
}
