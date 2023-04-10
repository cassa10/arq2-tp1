package exception

import "fmt"

type ProductNotFoundErr struct {
	Id   int64
	Name string
}

func (e ProductNotFoundErr) Error() string {
	if e.Id != 0 {
		return fmt.Sprintf("product with id %v not found", e.Id)
	}
	return fmt.Sprintf("product with name %v not found", e.Name)
}

type ProductCannotDelete struct {
	Id int64
}

func (e ProductCannotDelete) Error() string {
	return fmt.Sprintf("product with id %v cannot delete", e.Id)
}

type ProductCannotUpdate struct {
	Id int64
}

func (e ProductCannotUpdate) Error() string {
	return fmt.Sprintf("product with id %v cannot update", e.Id)
}

type ProductWithNoStock struct {
	Id int64
}

func (e ProductWithNoStock) Error() string {
	return fmt.Sprintf("product with id %v have no stock", e.Id)
}
