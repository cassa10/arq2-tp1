package model

import "fmt"

type CustomerAlreadyExistError struct {
	Email string
}

func (e CustomerAlreadyExistError) Error() string {
	return fmt.Sprintf("customer with email %s already exists", e.Email)
}

type CustomerNotFoundErr struct {
	Id int64
}

func (e CustomerNotFoundErr) Error() string {
	return fmt.Sprintf("customer with id %v not found", e.Id)
}

type CustomerCannotDelete struct {
	Id int64
}

func (e CustomerCannotDelete) Error() string {
	return fmt.Sprintf("customer with id %v cannot deleted", e.Id)
}

type CustomerCannotUpdate struct {
	Id int64
}

func (e CustomerCannotUpdate) Error() string {
	return fmt.Sprintf("customer with id %v cannot update", e.Id)
}
