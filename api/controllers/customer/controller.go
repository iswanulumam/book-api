package customer

import (
	"alta/book-api/api/common"
	"alta/book-api/models"
	"net/http"
	"strconv"

	echo "github.com/labstack/echo/v4"
)

type Controller struct {
	customerModel models.CustomerModel
}

func NewController(customerModel models.CustomerModel) *Controller {
	return &Controller{
		customerModel,
	}
}

func (controller *Controller) GetAll(c echo.Context) error {
	customer, err := controller.customerModel.Get()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
	}

	return c.JSON(http.StatusOK, customer)
}

func (controller *Controller) GetOne(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
	}

	user, errGetOne := controller.customerModel.GetOne(id)
	if errGetOne != nil {
		return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
	}

	type Response struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	response := Response{
		Name:  user.Name,
		Email: user.Email,
	}

	return c.JSON(http.StatusOK, response)
}

func (controller *Controller) PostOne(c echo.Context) error {
	// bind request value
	type Request struct {
		Name     string `json:"name" form:"name" validate:"required"`
		Email    string `json:"email" form:"email" validate:"required,email"`
		Password string `json:"password" form:"password" validate:"required"`
	}

	var userRequest Request
	if err := c.Bind(&userRequest); err != nil {
		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
	}

	user := models.Customer{
		Name:     userRequest.Name,
		Email:    userRequest.Email,
		Password: userRequest.Password,
	}

	_, err := controller.customerModel.Insert(user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
	}

	return c.JSON(http.StatusOK, common.NewSuccessOperationResponse())
}
