package handlers

import (
	"microservice/data"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// swagger:route DELETE /products/{id} products deleteProduct
// Deletes a product from the database
// Responses:
//  201: noContent

// DeleteProduct deletes a product from the database
func (p *Products) DeleteProduct(rw http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "unable to convert id string", http.StatusBadRequest)
	}

	p.logger.Println("handle DELETE Products for id:", id)

	err = data.DeleteProduct(id)
	if err == data.ErrProductNotFound {
		http.Error(rw, "product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "unknown error occurred", http.StatusInternalServerError)
		return
	}
}
