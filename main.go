package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/ele-fant/handlers"
	"github.com/nicholasjackson/env"
)

var bindAddress = env.String("BIND_ADDRESS", false, ":9090", "Bind address for the server")

func main() {

	//env.Parse()

	// register a function to DefaulServeMux
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	ph := handlers.NewProduct(l)

	sm := http.NewServeMux()
	sm.Handle("/", ph)

	// create a server
	s := &http.Server{
		Addr:    ":9090",
		Handler: sm,
		// control timeouts
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)

	// addr = bind (here to all IP-addesses with port 9090)
	// handler = if nil -> uses DefaulServeMux
	//http.ListenAndServe(":9090", sm)

	timeOutContext, _ := context.WithTimeout(context.Background(), 30*time.Second)

	s.Shutdown(timeOutContext)
}
