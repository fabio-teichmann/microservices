package main

import (
	"context"
	"log"
	"microservice/handlers"
	"net/http"
	"os"
	"os/signal"
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

	// this will block
	// server.ListenAndServe()
	// using a go routine instead to not block
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	// graceful shut-down:
	// waits for all running requests to be finished
	sigChan := make(chan os.Signal)
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
