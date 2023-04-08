package model

type UpdateCustomer struct {
	Lastname  string `json:"lastname" bson:"lastname" binding:"required"`
	Firstname string `json:"firstname" bson:"firstname" binding:"required"`
}
