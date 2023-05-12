package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "Nics",
		Price: 1.00,
		SKU:   "abs-abc-abc",
	}
	validator := NewValidator()
	err := validator.Validate(p)

	if err != nil {
		t.Fatal(err)
	}
}
