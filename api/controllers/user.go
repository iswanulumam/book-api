package controllers

import (
	"alta/book-api/models"
	"net/http"

	echo "github.com/labstack/echo/v4"
)

type UserController struct {
	userModel models.UserModel
}

func NewUserController(userModel models.UserModel) *UserController {
	return &UserController{
		userModel,
	}
}

func (controller *UserController) GetUser(c echo.Context) error {
	return c.String(http.StatusOK, "From Controller")
}
