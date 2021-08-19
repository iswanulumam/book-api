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
		Name     string `json:"name" form:"name"`
		Email    string `json:"email" form:"email"`
		Password string `json:"password" form:"password"`
	}

	var customerRequest Request
	if err := c.Bind(&customerRequest); err != nil {
		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
	}

	customer := models.Customer{
		Name:     customerRequest.Name,
		Email:    customerRequest.Email,
		Password: customerRequest.Password,
	}

	_, err := controller.customerModel.Insert(customer)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
	}

	return c.JSON(http.StatusOK, common.NewSuccessOperationResponse())
}

func (controller *Controller) EditOne(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
	}
	// bind request value
	type Request struct {
		Name     string `json:"name" form:"name"`
		Email    string `json:"email" form:"email"`
		Password string `json:"password" form:"password"`
	}

	var customerRequest Request
	if err := c.Bind(&customerRequest); err != nil {
		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
	}

	customer := models.Customer{
		Name:     customerRequest.Name,
		Email:    customerRequest.Email,
		Password: customerRequest.Password,
	}

	if _, err := controller.customerModel.Edit(customer, id); err != nil {
		return c.JSON(http.StatusNotFound, common.NewBadRequestResponse())
	}

	return c.JSON(http.StatusOK, common.NewSuccessOperationResponse())
}

func (controller *Controller) DeleteOne(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
	}

	if _, err := controller.customerModel.Delete(id); err != nil {
		return c.JSON(http.StatusInternalServerError, common.NewBadRequestResponse())
	}

	return c.JSON(http.StatusOK, common.NewSuccessOperationResponse())
}
