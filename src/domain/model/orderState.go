package model

type OrderState interface {
	//Confirm returns true when order mutates
	Confirm(order *Order) bool
	//Delivered returns true when order mutates
	Delivered(order *Order) bool
	String() string
}

type PendingOrderState struct{}

func (pS PendingOrderState) Confirm(order *Order) bool {
	order.State = ConfirmedOrderState{}
	return true
}

func (pS PendingOrderState) Delivered(_ *Order) bool {
	return false
}

func (pS PendingOrderState) String() string {
	return "PENDING"
}

type ConfirmedOrderState struct{}

func (cS ConfirmedOrderState) Confirm(_ *Order) bool {
	return false
}

func (cS ConfirmedOrderState) Delivered(order *Order) bool {
	order.State = DeliveredOrderState{}
	return true
}

func (cS ConfirmedOrderState) String() string {
	return "CONFIRMED"
}

type DeliveredOrderState struct{}

func (dS DeliveredOrderState) Confirm(_ *Order) bool {
	return false
}

func (dS DeliveredOrderState) Delivered(_ *Order) bool {
	return false
}

func (dS DeliveredOrderState) String() string {
	return "DELIVERED"
}
