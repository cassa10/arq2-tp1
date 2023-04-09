package dto

import "github.com/cassa10/arq2-tp1/src/domain/model"

type ProductCreateReq struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Price       float64 `json:"price" binding:"required,min=0"`
	Category    string  `json:"category" binding:"required"`
	Stock       int     `json:"stock" binding:"required,min=1"`
}

func (req *ProductCreateReq) MapToModel(sellerId int64) model.Product {
	return model.Product{
		SellerId:    sellerId,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Category:    req.Category,
		Stock:       req.Stock,
	}
}
