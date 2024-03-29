package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

// Product defines the structure for an API product
// swagger:model
type Product struct {
	// product id
	//
	// required: true
	// min: 1
	ID int `json:"id"`

	// product name
	//
	// required: true
	// max: 90
	Name string `json:"name" validate:"required"`

	// description of the product
	//
	// required: false
	// max: 120
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku" validate:"required,sku"` // custom validation function
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

type Products []*Product

var productList = []*Product{
	{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy, milky coffee",
		Price:       2.45,
		SKU:         "abc123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	{
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

func UpdateProduct(id int, p *Product) error {
	_, pos, err := getProductById(id)
	if err != nil {
		return err
	}
	p.ID = id
	productList[pos] = p

	return nil
}

func DeleteProduct(id int) error {
	_, pos, err := getProductById(id)
	if err != nil {
		return err
	}
	productList = append(productList[:pos], productList[pos+1:]...)
	return nil
}

func getNextId() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
}

// structured error
var ErrProductNotFound = fmt.Errorf("Product not found")

func getProductById(id int) (*Product, int, error) {
	for i, prod := range productList {
		if prod.ID == id {
			return prod, i, nil
		}
	}

	return nil, -1, ErrProductNotFound
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
