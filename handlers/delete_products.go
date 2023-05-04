package handlers

import (
	"microservice/data"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

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
