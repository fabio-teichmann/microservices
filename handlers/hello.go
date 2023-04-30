package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	// dependency injection
	logger *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

// implements HTTP Handler Interface
func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.logger.Println("Hello World")
	// read slices of data from the http request's body
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// inform sender on error
		http.Error(rw, "Oops", http.StatusBadRequest)
		return
	}

	// log.Printf("Data: %s\n", data)
	// print back to the response writer
	fmt.Fprintf(rw, "Hello %s", data)
}
