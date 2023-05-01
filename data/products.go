package data

import (
	"encoding/json"
	"io"
	"time"
)

// Product defines the structure for an API product
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

type Products []*Product

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy, milky coffee",
		Price:       2.45,
		SKU:         "abc123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Strong coffee",
		Price:       1.50,
		SKU:         "def334",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}

func GetProducts() Products {
	return productList
}

func AddProduct(p *Product) {
	p.ID = getNextId()
	productList = append(productList, p)
}

func getNextId() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
}

func (p *Products) ToJSON(w io.Writer) error {
	// encapsulate json translation logic
	// Encoder does not allocate additional memory (buffering) but rather
	// writes it directly to stream. This reduces memory and overhead of
	// the service.
	encoder := json.NewEncoder(w)
	return encoder.Encode(p)
}

func (p *Product) FromJSON(r io.Reader) error {
	// encapsulate json translation logic
	// Decoder translates back from JSON into struct
	decoder := json.NewDecoder(r)
	return decoder.Decode(p)
}
