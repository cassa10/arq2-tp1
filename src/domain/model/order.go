package model

import (
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
	FindById(id int64) (Order, error)
	Create(order Order) (Order, error)
	Update(order Order) (Order, error)
	Delete(id int64) (bool, error)
}
