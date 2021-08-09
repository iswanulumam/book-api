package user

import (
	"alta/book-api/api/common"
	"alta/book-api/api/controllers/user/request"
	"alta/book-api/api/controllers/user/response"
	"alta/book-api/models"
	"net/http"

	echo "github.com/labstack/echo/v4"
)

type Controller struct {
	userModel models.UserModel
}

func NewController(userModel models.UserModel) *Controller {
	return &Controller{
		userModel,
	}
}

func (controller *Controller) GetUser(c echo.Context) error {
	users, err := controller.userModel.Get()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
	}

	result := response.NewGetUserResponse(users)

	return c.JSON(http.StatusOK, result)
}

func (controller *Controller) PostUser(c echo.Context) error {
	// bind request value
	var userRequest request.PostUserRequest
	if err := c.Bind(&userRequest); err != nil {
		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
	}

	// validation input
	postUserValidation := userRequest.PostUserValidation()

	if postUserValidation != nil {
		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
	}

	// bind userReqeust to user modes
	user := models.User{
		Name:     userRequest.Name,
		Email:    userRequest.Email,
		Password: userRequest.Password,
	}

	userResult, err := controller.userModel.Insert(user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
	}

	result := response.NewPostUserResponse(userResult)

	return c.JSON(http.StatusOK, result)
}
