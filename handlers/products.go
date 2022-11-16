package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/ele-fant/data"
	"github.com/gorilla/mux"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

/* DEPRECATED
func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	// handle an insert == POST
	if r.Method == http.MethodPost {
		p.addProducts(rw, r)
		return
	}

	// handle an update == PUT
	if r.Method == http.MethodPut {
		//p.l.Println("Handle PUT Product")
		// expect ID in the URI
		reg := regexp.MustCompile(`/([0-9]+)`)
		group := reg.FindAllStringSubmatch(r.URL.Path, -1)

		if len(group) != 1 {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		if len(group[0]) != 2 {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		idString := group[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
		}

		p.l.Println("Got ID", id)

		p.updateProduct(id, rw, r)
		return
	}

	// catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}
*/

// get Products from the data store
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Product")

	lp := data.GetProducts()

	// serialize the list to json
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) AddProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")

	prod := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&prod)
}

func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // extracts params as map
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "unable to convert id to int", http.StatusInternalServerError)
	}

	p.l.Println("Handle PUT Product", id)

	prod := r.Context().Value(KeyProduct{}).(data.Product)

	err = data.UpdateProduct(id, &prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, err.Error(), http.StatusNotFound)
	}
	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}

/*
	func (p Products) UpdateProducts(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(rw, "Unable to convert id", http.StatusBadRequest)
			return
		}

		p.l.Println("Handle PUT Product", id)
		prod := r.Context().Value(KeyProduct{}).(data.Product)

		err = data.UpdateProduct(id, &prod)
		if err == data.ErrProductNotFound {
			http.Error(rw, "Product not found", http.StatusNotFound)
			return
		}

		if err != nil {
			http.Error(rw, "Product not found", http.StatusInternalServerError)
			return
		}
	}
*/
type KeyProduct struct{}

func (p Products) MiddlewareValidateProduct(next http.Handler) http.Handler {

	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := data.Product{}

		err := prod.FromJSON(r.Body)
		if err != nil {
			p.l.Panicln("[ERROR] deserializing product", err)
			http.Error(rw, "unable to unmarshal json", http.StatusBadRequest)
			return
		}

		// validate product
		err = prod.Validate()
		if err != nil {
			p.l.Panicln("[ERROR] validating product", err)
			http.Error(
				rw,
				fmt.Sprintf("Error: validating product: %s", err),
				http.StatusBadRequest,
			)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)

		next.ServeHTTP(rw, r)
	})
}
