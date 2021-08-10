package api

import (
	"alta/book-api/api/controllers/user"

	echo "github.com/labstack/echo/v4"
)

func RegisterPath(e *echo.Echo, userController *user.Controller) {
	e.GET("/users", userController.GetUser)
	e.GET("/users/:id", userController.GetUserOne)
	e.POST("/users", userController.PostUser)
}
