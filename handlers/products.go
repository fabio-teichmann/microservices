// Package classification of Product API
//
// Documentation for Product API
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package handlers

import (
	"context"
	"fmt"
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

// func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
// 	if r.Method == http.MethodGet {
// 		p.getProducts(rw, r)
// 		return
// 	}

// 	if r.Method == http.MethodPost {
// 		p.addProduct(rw, r)
// 		return
// 	}
// 	// handle UPDATE
// 	if r.Method == http.MethodPut {
// 		// expect ID in the URI
// 		reg := regexp.MustCompile(`/([0-9]+)`)
// 		groups := reg.FindAllStringSubmatch(r.URL.Path, -1)

// 		if len(groups) != 1 {
// 			http.Error(rw, "invalid URI", http.StatusBadRequest)
// 			return
// 		}

// 		if len(groups[0]) != 2 {
// 			http.Error(rw, "incalid URI", http.StatusBadRequest)
// 		}

// 		idString := groups[0][1]
// 		id, err := strconv.Atoi(idString)
// 		if err != nil {
// 			http.Error(rw, "invalid URI", http.StatusBadRequest)
// 			return
// 		}
// 		// p.logger.Println("got id:", id)

// 		p.updateProduct(id, rw, r)
// 		return
// 	}

// 	// catch all other cases
// 	rw.WriteHeader(http.StatusMethodNotAllowed)
// }

// A list of products returned in the response
// swagger:response productsResponse
type productsResponse struct {
	// All products in the system
	// in: body
	Body []data.Product
}

type KeyProduct struct{}

func (p Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := data.Product{}

		err := prod.FromJSON(r.Body)
		if err != nil {
			http.Error(rw, "unable to unmarshal json", http.StatusBadRequest)
			return
		}

		// validate the product
		err = prod.Validate()
		if err != nil {
			p.logger.Println("[ERROR] validating product", err)
			http.Error(rw, fmt.Sprintf("Error validating product: %s\n", err), http.StatusBadRequest)
			return
		}
		// store passed data in context to be accesible for handling
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)

		next.ServeHTTP(rw, r)
	})
}
