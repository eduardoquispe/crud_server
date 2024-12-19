package dtos

import "github.com/go-playground/validator/v10"

type UserRequest struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstName" validate:"required,min=2,max=50"`
	LastName  string `json:"lastName" validate:"required,min=2,max=50"`
	Email     string `json:"email" validate:"required,email"`
	Age       int    `json:"age" validate:"required,min=18,max=99"`
}

var validate = validator.New()

func (r *UserRequest) Validate() error {
	return validate.Struct(r)
}