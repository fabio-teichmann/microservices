package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// HandleFunc is a convenience function that registers a function on a path called
	// "Default ServMux" (=http handler / http request multiplexer)
	// The ServMux determines which function gets activated.
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello World")
		// read slices of data from the http request's body
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			// inform sender on error
			http.Error(rw, "Oops", http.StatusBadRequest)
			return
			// log.Println(err)
		}

		// log.Printf("Data: %s\n", data)
		// print back to the response writer
		fmt.Fprintf(rw, "Hello %s", data)
	})

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Good bye World")
	})

	// BIND address, HANDLER
	// most basic webserver; if HANDLER not specified, it uses the default ServMux
	http.ListenAndServe(":9090", nil)
}
