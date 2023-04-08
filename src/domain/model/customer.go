package model

import (
	"context"
	"github.com/cassa10/arq2-tp1/src/domain/util"
)

type Customer struct {
	Id        int64  `json:"id" bson:"_id"`
	Firstname string `json:"firstname" bson:"firstname" binding:"required"`
	Lastname  string `json:"lastname" bson:"lastname" binding:"required"`
	Email     string `json:"email" bson:"email" binding:"required,email"`
}

func (c *Customer) Merge(updateCustomer UpdateCustomer) {
	c.Firstname = updateCustomer.Firstname
	c.Lastname = updateCustomer.Lastname
}

func (c *Customer) String() string {
	return util.ParseStruct("Customer", c)
}

type CustomerRepository interface {
	FindById(ctx context.Context, id int64) (*Customer, error)
	Create(ctx context.Context, customer Customer) (*Customer, error)
	Update(ctx context.Context, customer Customer) (*Customer, error)
	Delete(ctx context.Context, id int64) (bool, error)
}
