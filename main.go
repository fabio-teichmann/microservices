package main

import (
	"log"
	"microservice/handlers"
	"net/http"
	"os"
	"time"
)

func main() {
	logger := log.New(os.Stdout, "product-api", log.LstdFlags)
	handler := handlers.NewHello(logger)
	gbHandler := handlers.NewGoodbye(logger)

	serveMux := http.NewServeMux()
	serveMux.Handle("/", handler)
	serveMux.Handle("/goodbye", gbHandler)

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
		Addr:         ":9090",
		Handler:      serveMux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	server.ListenAndServe()
	// BIND address, HANDLER
	// most basic webserver; if HANDLER not specified, it uses the default ServMux
	// http.ListenAndServe(":9090", serveMux)
}
