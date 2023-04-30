package main

import (
	"log"
	"microservice/handlers"
	"net/http"
	"os"
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

	// http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
	// 	log.Println("Good bye World")
	// })

	// BIND address, HANDLER
	// most basic webserver; if HANDLER not specified, it uses the default ServMux
	http.ListenAndServe(":9090", serveMux)
}
