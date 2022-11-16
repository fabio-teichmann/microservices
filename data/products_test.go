package data

import (
	"testing"
)

func TestValidate(t *testing.T) {
	prod := &Product{}
	prod.Name = "tea"
	prod.Price = 1.0
	prod.SKU = "abs-aaas-aaas"

	err := prod.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
