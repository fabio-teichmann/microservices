package handlers

import (
	"context"
	"log"
	"microservice/data"
	"net/http"
)

type Products struct {
	logger    *log.Logger
	validator *data.Validation
}

func NewProducts(l *log.Logger, v *data.Validation) *Products {
	return &Products{l, v}
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

type KeyProduct struct{}

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages []string `json:"messages"`
}

func (p Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")

		prod := data.Product{}

		err := prod.FromJSON(r.Body)
		if err != nil {
			http.Error(rw, "unable to unmarshal json", http.StatusBadRequest)
			return
		}

		// validate the product
		errs := p.validator.Validate(prod)
		if errs != nil {
			p.logger.Println("[ERROR] validating product", errs)

			// return the validation messages as an array
			rw.WriteHeader(http.StatusUnprocessableEntity)
			data.ToJSON(&ValidationError{Messages: errs.Errors()}, rw)
			// http.Error(rw, fmt.Sprintf("Error validating product: %s\n", err), http.StatusBadRequest)
			return
		}
		// store passed data in context to be accesible for handling
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)

		next.ServeHTTP(rw, r)
	})
}
