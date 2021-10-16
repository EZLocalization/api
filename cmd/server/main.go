package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/ez-api/internal/app/server"
)

const (
	applicationName = "Admin API Server"
)

// StartServer server running
func StartServer(configPath string) {
	ctx := context.Background()
	// ctx = context.WithValue(ctx, constant.ContextApplication, applicationName)

	// load configuration
	cfg, err := GetConfigruation(ctx, configPath)
	if err != nil {
		// logger.Fatal(err)
	}

	// set logging configuration
	// log.SetGlobalConfiguration(cfg.LoggingConfig)

	// create http server
	server, err := server.New(ctx, cfg.ServerConfig)
	if err != nil {
		// logger.Fatal(err)
	}

	// http start start
	server.Start(ctx)
	defer server.Stop(ctx)

	// wait signal
	WaitSignal(ctx, syscall.SIGKILL, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
}

// WaitSignal this is blocking function until signal comes or context cancel
// ex) signals: syscall.SIGKILL, syscall.SIGTERM, syscall.SIGINT, os.Interrupt
func WaitSignal(ctx context.Context, signals ...os.Signal) {

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, signals...)
	sig := <-sigs
	fmt.Printf("signal: %s", sig.String())
	signal.Stop(sigs)
}

func main() {
	// Run Command
	cmd := NewCommandLine()
	cmd.Execute()
}
