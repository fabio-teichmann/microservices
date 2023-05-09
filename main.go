package main

import (
	"context"
	"log"
	"microservice/data"
	"microservice/handlers"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"github.com/nicholasjackson/env"
)

var bindAddress = env.String("BIND_ADDRESS", false, ":9090", "Bind address for the server")

func main() {
	env.Parse()

	logger := log.New(os.Stdout, "products-api", log.LstdFlags)
	validator := data.NewValidator()

	productHandler := handlers.NewProducts(logger, validator)

	// serveMux := http.NewServeMux()
	// replacing serveMux with Gorilla mux router
	serveMux := mux.NewRouter()
	// create subrouter for GET requests
	getRouter := serveMux.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", productHandler.GetProducts)

	putRouter := serveMux.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", productHandler.UpdateProduct)
	// register middleware on subrouter
	putRouter.Use(productHandler.MiddlewareProductValidation)

	postRouter := serveMux.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", productHandler.AddProduct)
	// register middleware on subrouter
	postRouter.Use(productHandler.MiddlewareProductValidation)

	deleteRouter := serveMux.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/{id:[0-9]+}", productHandler.DeleteProduct)

	// documentation handler using ReDoc middleware
	options := middleware.RedocOpts{SpecURL: "/swagger.yaml"} // per defaults looks for .json
	sh := middleware.Redoc(options, nil)

	getRouter.Handle("/docs", sh)
	// serve swagger.yaml
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

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
