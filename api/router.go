package api

import (
	"alta/book-api/api/controllers/user"

	echo "github.com/labstack/echo/v4"
)

func RegisterPath(e *echo.Echo, userController *user.Controller) {

	e.GET("/users", userController.GetUser)
	e.POST("/users", userController.PostUser)

	//health check
	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(200)
	})
}
