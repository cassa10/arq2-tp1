package query

import (
	"context"
	"github.com/cassa10/arq2-tp1/src/domain/model"
)

type SearchProduct struct {
	productRepo model.ProductRepository
}

func NewSearchProduct(productRepo model.ProductRepository) *SearchProduct {
	return &SearchProduct{
		productRepo: productRepo,
	}
}

func (q SearchProduct) Do(ctx context.Context, filters model.ProductSearchFilter, pagingReq model.PagingRequest) ([]model.Product, model.Paging, error) {
	return q.productRepo.Search(ctx, filters, pagingReq)
}
