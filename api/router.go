package api

import (
	"alta/book-api/api/controllers/customer"

	echo "github.com/labstack/echo/v4"
)

func RegisterPath(e *echo.Echo, customerController *customer.Controller) {
	// ------------------------------------------------------------------
	// Login & register
	// ------------------------------------------------------------------
	e.POST("/customers/register", customerController.PostCustomerController)
	e.POST("/customers/login", customerController.DeleteCustomerController)

	// ------------------------------------------------------------------
	// CRUD Customer
	// ------------------------------------------------------------------
	e.GET("/customers", customerController.GetAllCustomerController)
	e.GET("/customers/:id", customerController.GetCustomerController)
	e.PUT("/customers/:id", customerController.EditCustomerController)
	e.DELETE("/customers/:id", customerController.DeleteCustomerController)
}
