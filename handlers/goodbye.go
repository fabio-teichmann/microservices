package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Bye struct {
	l *log.Logger
}

func NewBye(l *log.Logger) *Bye {
	return &Bye{l}
}

func (b *Bye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	d, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "whoopsi", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "Goodbye %s!", d)
}
