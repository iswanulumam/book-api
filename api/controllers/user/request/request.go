package request

import (
	validator "github.com/go-playground/validator/v10"
)

type PostUserRequest struct {
	Name     string `json:"name" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required"`
}

func (p PostUserRequest) PostUserValidation() error {
	validate := validator.New()

	err := validate.Struct(p)
	if err == nil {
		return nil
	}

	return err
}
