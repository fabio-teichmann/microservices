package handlers

import (
	"log"
	"microservice/data"
	"net/http"
)

type Products struct {
	logger *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	// handle UPDATE

	// catch all other cases
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	listProducts := data.GetProducts()
	// encoding JSON using marshal -> allocates data to memory
	// d, err := json.Marshal(listProducts) DEPRECATED
	err := listProducts.ToJSON(rw)
	if err != nil {
		http.Error(rw, "unable to marshal JSON", http.StatusInternalServerError)
	}
	// write marshalled data
	// rw.Write(d) DEPRECATED
}
