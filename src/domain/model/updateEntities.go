package model

type UpdateCustomer struct {
	Lastname  string `json:"lastname" bson:"lastname" binding:"required"`
	Firstname string `json:"firstname" bson:"firstname" binding:"required"`
	Email     string `json:"email" bson:"email" binding:"required,email"`
}

type UpdateSeller struct {
	Name  string `json:"name" bson:"name" binding:"required"`
	Email string `json:"email" bson:"email" binding:"required,email"`
}
