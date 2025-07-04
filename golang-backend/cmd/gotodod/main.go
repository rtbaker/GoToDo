package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Setup OS signal handling
	//
	// All http servers are started using a goroutine so they
	// return straight away. We then sit and wait on an os signal
	// to bring things to a halt.
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	go func() { <-sigs; cancel() }()

	// Create our app
	app := NewApplication()

	if err := app.LoadConfig(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Run our app
	if err := app.Run(ctx); err != nil {
		app.Close()
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Wait for OS signal to end our server.
	<-ctx.Done()

	// Clean up program.
	if err := app.Close(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
