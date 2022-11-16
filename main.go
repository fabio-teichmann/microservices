package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/ele-fant/handlers"
	"github.com/gorilla/mux"
	"github.com/nicholasjackson/env"
)

var bindAddress = env.String("BIND_ADDRESS", false, ":9090", "Bind address for the server")

func main() {

	env.Parse()

	// register a function to DefaulServeMux
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	// create the handlers
	prod_handler := handlers.NewProducts(l)

	// create a new serve mux and register the handlers
	sm := mux.NewRouter()

	//.Methods("GET") --> can specify valid methods
	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", prod_handler.GetProducts)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", prod_handler.UpdateProduct)
	putRouter.Use(prod_handler.MiddlewareValidateProduct)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", prod_handler.AddProducts)
	postRouter.Use(prod_handler.MiddlewareValidateProduct)

	//sm.Handle("/products", ph)

	// create a server
	s := &http.Server{
		Addr:         *bindAddress,      // configure the bind address
		Handler:      sm,                // set the default handler
		ErrorLog:     l,                 // set the logger for the server
		IdleTimeout:  120 * time.Second, // max time to read request from client
		ReadTimeout:  1 * time.Second,   // max time to write request to client
		WriteTimeout: 1 * time.Second,   // max time for connection using TCP Keep-Alive
	}

	// start the server
	go func() {
		l.Println("Starting server on port 9090")

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	// Block until a signal is received.
	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)

	// addr = bind (here to all IP-addesses with port 9090)
	// handler = if nil -> uses DefaulServeMux
	//http.ListenAndServe(":9090", sm)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	timeOutContext, _ := context.WithTimeout(context.Background(), 30*time.Second)

	s.Shutdown(timeOutContext)
}
