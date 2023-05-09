package handlers

import (
	"microservice/data"
	"net/http"
)

// swagger:route GET /products products getProducts
// Returns a list of products
//
//  Consumes:
//   - application/json
//
//  Deprecated: false
//
//  Responses:
//   200: productsResponse
//	 500: unableToMarshalJSON

// GetProducts returns the products from the data store
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {

	p.logger.Println("handle GET Products")

	rw.Header().Add("Content-Type", "application/json")

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
