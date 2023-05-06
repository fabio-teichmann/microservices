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

import "microservice/data"

// A list of products returned in the response
// swagger:response productsResponse
type productsResponse struct {
	// All products in the system
	// in: body
	Body data.Product
}

// swagger:response noContent
type productsNoContent struct {
}

//swagger:response unableToMarshalJSON
type productsUnableToMarshalJSON struct {
}

// swagger:parameters deleteProduct
type productIDParameterWrapper struct {
	// The id of the product to delete from the database
	// in: path
	// required: true
	ID int `json:"id"`
}
