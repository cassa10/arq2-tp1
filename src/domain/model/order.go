package model

import (
	"context"
	"github.com/cassa10/arq2-tp1/src/domain/util"
	"time"
)

type Order struct {
	Id              int64      `json:"id"`
	CreatedOn       time.Time  `json:"createdOn"`
	UpdatedOn       time.Time  `json:"updatedOn"`
	DeliveryDate    time.Time  `json:"deliveryDate"`
	State           OrderState `json:"state"`
	Product         Product    `json:"product"`
	Customer        Customer   `json:"customer"`
	DeliveryAddress Address    `json:"deliveryAddress"`
}

func (o *Order) String() string {
	return util.ParseStruct("Order", o)
}

// Confirm returns true when order mutates
func (o *Order) Confirm() bool {
	return o.State.Confirm(o)
}

// Delivered returns true when order mutates
func (o *Order) Delivered() bool {
	return o.State.Delivered(o)
}

type OrderRepository interface {
	FindById(ctx context.Context, id int64) (*Order, error)
	Create(ctx context.Context, order Order) (int64, error)
	Update(ctx context.Context, order Order) (bool, error)
	Delete(ctx context.Context, id int64) (bool, error)
}
