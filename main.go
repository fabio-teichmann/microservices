package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	// register a function to DefaulServeMux
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {

	})

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Goodbye World!")
	})
	// addr = bind (here to all IP-addesses with port 9090)
	// handler = if nil -> uses DefaulServeMux
	http.ListenAndServe(":9090", nil)
}
