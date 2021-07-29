package main

import (
	"alta/book-api/api"
	"alta/book-api/api/controllers"
	"alta/book-api/config"
	"alta/book-api/models"
	"alta/book-api/util"

	"fmt"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	//load config if available or set to default
	config := config.GetConfig()

	//initialize database connection based on given config
	db := util.MysqlDatabaseConnection(config)

	//initiate user model
	userModel := models.NewUserModel(db)

	//initiate user controller
	userController := controllers.NewUserController(userModel)

	//create echo http
	e := echo.New()

	//register API path and controller
	api.UserRegisterPath(e, userController)

	// run server
	address := fmt.Sprintf("localhost:%d", config.Port)

	if err := e.Start(address); err != nil {
		log.Info("shutting down the server")
	}
}
