package handlers

import (
	"microservice/data"
	"net/http"
)

func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.logger.Println("handle GET Products")

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
