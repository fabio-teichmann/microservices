package handlers

import (
	"microservice/data"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	// inputs extracted into mux.Vars()
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "unable to convert id string", http.StatusBadRequest)
	}

	p.logger.Println("handle PUT Products", id)
	// cast KeyProduct into target structure
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	err = data.UpdateProduct(id, &prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "product not found", http.StatusInternalServerError)
		return
	}
}
