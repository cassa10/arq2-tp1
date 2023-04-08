package query

import (
	"context"
	"github.com/cassa10/arq2-tp1/src/domain/model"
)

type FindCustomer struct {
	customerRepo model.CustomerRepository
}

func NewFindCustomer(customerRepo model.CustomerRepository) *FindCustomer {
	return &FindCustomer{
		customerRepo: customerRepo,
	}
}

func (q FindCustomer) Do(ctx context.Context, id int64) (*model.Customer, error) {
	return q.customerRepo.FindById(ctx, id)
}
