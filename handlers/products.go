package handlers

import (
	"log"
	"microservice/data"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.logger.Println("handle POST Products")

	prod := &data.Product{}
	// pass reader into the FromJSON function
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "unable to unmarshal json", http.StatusBadRequest)
	}
	data.AddProduct(prod)
	// p.logger.Printf("Prod %#v", prod)
}

func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	// inputs extracted into mux.Vars()
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "unable to convert id string", http.StatusBadRequest)
	}

	p.logger.Println("handle PUT Products", id)

	prod := &data.Product{}

	err = prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "unable to unmarshal json", http.StatusBadRequest)
		return
	}
	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "product not found", http.StatusInternalServerError)
		return
	}
}
