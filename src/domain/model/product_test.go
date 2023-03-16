package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Product_String(t *testing.T) {
	product1 := Product{
		Name:        "Jabon",
		Description: "Un jabon lindo",
	}
	product2 := Product{
		Name:        "Galletitas",
		Description: "Galletitas sonrisa",
	}
	assert.Equal(t, `[Product]{"Name":"Jabon","Description":"Un jabon lindo"}`, product1.String())
	assert.Equal(t, `[Product]{"Name":"Galletitas","Description":"Galletitas sonrisa"}`, product2.String())
}
