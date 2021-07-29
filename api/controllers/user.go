package controllers

import (
	"alta/book-api/api/common"
	"alta/book-api/api/controllers/response"
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
	users, err := controller.userModel.Get()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
	}

	response := response.NewGetUserResponse(users)

	return c.JSON(http.StatusOK, response)
}

func (controller *UserController) PostUser(c echo.Context) error {
	u := new(models.User)

	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
	}

	user := models.User{
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
	}

	userResult, err := controller.userModel.Insert(user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
	}

	response := response.NewPostUserResponse(userResult)

	return c.JSON(http.StatusOK, response)
}
