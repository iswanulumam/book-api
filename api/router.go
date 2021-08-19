package api

import (
	"alta/book-api/api/controllers/customer"

	echo "github.com/labstack/echo/v4"
)

func RegisterPath(e *echo.Echo, customerController *customer.Controller) {
	e.GET("/customers", customerController.GetAll)
	e.POST("/customers", customerController.PostOne)
	e.GET("/customers/:id", customerController.GetOne)
}
