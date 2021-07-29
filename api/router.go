package api

import (
	"alta/book-api/api/controllers"

	echo "github.com/labstack/echo/v4"
)

func UserRegisterPath(e *echo.Echo, c *controllers.UserController) {

	e.GET("/users", c.GetUser)

	//health check
	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(200)
	})
}
