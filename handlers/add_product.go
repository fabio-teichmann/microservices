package handlers

import (
	"microservice/data"
	"net/http"
)

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.logger.Println("handle POST Products")

	prod := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&prod)
}
