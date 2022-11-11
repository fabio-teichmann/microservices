package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ele-fant/microservices/handlers"
)

func main() {
	// register a function to DefaulServeMux
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	// addr = bind (here to all IP-addesses with port 9090)
	// handler = if nil -> uses DefaulServeMux
	http.ListenAndServe(":9090", nil)
}
