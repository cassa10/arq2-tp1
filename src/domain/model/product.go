package model

import (
	"github.com/cassa10/arq2-tp1/src/domain/util"
)

type Product struct {
	Id          int64   `json:"id"`
	SellerId    int64   `json:"sellerId"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
	Stock       int     `json:"stock"`
}

func (p Product) String() string {
	return util.ParseStruct("Product", p)
}

type ProductRepository interface {
	AddProduct(sellerId int64, product Product) (bool, error)
	UpdateProduct(sellerId int64, product Product) (bool, error)
	DeleteProduct(sellerId int64, productId int64) (bool, error)
	Search(filters ProductSearchFilter) ([]Product, error)
}
