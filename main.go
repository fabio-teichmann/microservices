package main

import (
	"context"
	"log"
	"microservice/handlers"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/nicholasjackson/env"
)

var bindAddress = env.String("BIND_ADDRESS", false, ":9090", "Bind address for the server")

func main() {
	env.Parse()

	logger := log.New(os.Stdout, "product-api", log.LstdFlags)

	productHandler := handlers.NewProducts(logger)

	serveMux := http.NewServeMux()
	serveMux.Handle("/", productHandler)

	// HandleFunc is a convenience function that registers a function on a path called
	// "Default ServMux" (=http handler / http request multiplexer)
	// The ServMux determines which function gets activated.
	// http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {

	// })

	// creating an HTTP server
	// Some of the properties are:
	// - address
	// - handler
	// - TLSConfig
	// - read timeout (max time reading from client)
	// - write timeout
	// - idle timeout (keep connections alive)
	// These parameters should be tuned based on the needs
	server := &http.Server{
		// Addr:         ":9090",
		Addr:         *bindAddress,
		Handler:      serveMux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// this will block
	// server.ListenAndServe()
	// using a go routine instead to not block
	go func() {
		logger.Println("Starting server on port 9090")

		err := server.ListenAndServe()
		if err != nil {
			logger.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// graceful shut-down:
	// waits for all running requests to be finished
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	logger.Println("Received terminate, graceful shutdown", sig)

	timeoutContext, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(timeoutContext)

	// BIND address, HANDLER
	// most basic webserver; if HANDLER not specified, it uses the default ServMux
	// http.ListenAndServe(":9090", serveMux)
}
