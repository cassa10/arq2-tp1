package model

import (
	"encoding/json"
	"strings"
)

const (
	pendingState   = "PENDING"
	confirmedState = "CONFIRMED"
	deliveredState = "DELIVERED"
)

var stateMapper = map[string]OrderState{
	pendingState:   PendingOrderState{},
	confirmedState: ConfirmedOrderState{},
	deliveredState: DeliveredOrderState{},
}

type OrderState interface {
	// Confirm returns true when order mutates
	Confirm(order *Order) bool
	// Delivered returns true when order mutates
	Delivered(order *Order) bool
	String() string
	UnmarshalJSON(b []byte) error
	MarshalJSON() ([]byte, error)
}

func unmarshalJSONOrderState(orderState OrderState, b []byte) error {
	state := orderState.String()
	if err := json.Unmarshal(b, &state); err != nil {
		return err
	}
	return nil
}

func marshalJSONOrderState(orderState OrderState) ([]byte, error) {
	return json.Marshal(orderState.String())
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
	return pendingState
}

func (pS PendingOrderState) UnmarshalJSON(b []byte) error {
	return unmarshalJSONOrderState(pS, b)
}

func (pS PendingOrderState) MarshalJSON() ([]byte, error) {
	return marshalJSONOrderState(pS)
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
	return confirmedState
}

func (cS ConfirmedOrderState) UnmarshalJSON(b []byte) error {
	return unmarshalJSONOrderState(cS, b)
}

func (cS ConfirmedOrderState) MarshalJSON() ([]byte, error) {
	return marshalJSONOrderState(cS)
}

type DeliveredOrderState struct{}

func (dS DeliveredOrderState) Confirm(_ *Order) bool {
	return false
}

func (dS DeliveredOrderState) Delivered(_ *Order) bool {
	return false
}

func (dS DeliveredOrderState) String() string {
	return deliveredState
}

func (dS DeliveredOrderState) UnmarshalJSON(b []byte) error {
	return unmarshalJSONOrderState(dS, b)
}

func (dS DeliveredOrderState) MarshalJSON() ([]byte, error) {
	return marshalJSONOrderState(dS)
}

func GetStateByString(state string) (OrderState, bool) {
	orderState, ok := stateMapper[strings.ToUpper(state)]
	return orderState, ok
}
