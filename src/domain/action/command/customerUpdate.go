package command

import (
	"context"
	"github.com/cassa10/arq2-tp1/src/domain/action/query"
	"github.com/cassa10/arq2-tp1/src/domain/model"
)

type UpdateCustomer struct {
	customerRepo      model.CustomerRepository
	findCustomerQuery query.FindCustomer
}

func NewUpdateCustomer(customerRepo model.CustomerRepository, findCustomer query.FindCustomer) *UpdateCustomer {
	return &UpdateCustomer{
		customerRepo:      customerRepo,
		findCustomerQuery: findCustomer,
	}
}

func (c UpdateCustomer) Do(ctx context.Context, customerId int64, updateCustomer model.UpdateCustomer) error {
	customer, err := c.findCustomerQuery.Do(ctx, customerId)
	if err != nil {
		return err
	}
	customer.Merge(updateCustomer)
	if _, err := c.customerRepo.Update(ctx, *customer); err != nil {
		return err
	}
	return nil
}
