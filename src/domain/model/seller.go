package model

import "github.com/cassa10/arq2-tp1/src/domain/util"

type Seller struct {
	Id       int64     `json:"id"`
	Name     string    `json:"name"`
	Products []Product `json:"products"`
}

func (s *Seller) String() string {
	return util.ParseStruct("Seller", s)
}

type SellerRepository interface {
	FindById(id int64) (Seller, error)
	Create(seller Seller) (Seller, error)
	Update(seller Seller) (Seller, error)
	Delete(id int64) (bool, error)
}
