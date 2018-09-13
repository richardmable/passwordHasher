package main

import (
	"context"
	"fmt"
	"net/http"
)

func gracefulShutdown(srv *http.Server, idleConnsClosed chan struct{}, sigStop chan bool) {
	// block until we receive the shutdown signal
	<-sigStop
	fmt.Println("Starting shutdown...finishing open processes")
	// received an interrupt signal, shut down
	if err := srv.Shutdown(context.Background()); err != nil {
		// error from closing listeners, or context timeout
		fmt.Printf("HTTP server Shutdown: %v", err)
	}
	close(idleConnsClosed)
	fmt.Println("Shutdown complete.")
}
