package api

import (
	"alta/book-api/api/controllers/customer"

	echo "github.com/labstack/echo/v4"
)

func RegisterPath(e *echo.Echo, customerController *customer.Controller) {
	e.POST("/customers/register", customerController.PostOne)
	/// e.POST("/customers/login", customerController.PostLogin)

	e.GET("/customers", customerController.GetAll)
	e.GET("/customers/:id", customerController.GetOne)
	e.PUT("/customers/:id", customerController.EditOne)
	e.DELETE("/customers/:id", customerController.DeleteOne)
}
