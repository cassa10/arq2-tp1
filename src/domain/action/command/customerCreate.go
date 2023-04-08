package command

import (
	"context"
	"github.com/cassa10/arq2-tp1/src/domain/model"
)

type CreateCustomer struct {
	customerRepo model.CustomerRepository
}

func NewCreateCustomer(customerRepo model.CustomerRepository) *CreateCustomer {
	return &CreateCustomer{
		customerRepo: customerRepo,
	}
}

func (c CreateCustomer) Do(ctx context.Context, customer model.Customer) error {
	if _, err := c.customerRepo.Create(ctx, customer); err != nil {
		return err
	}
	return nil
}
