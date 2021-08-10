package user

import (
	"alta/book-api/api/controllers/user/response"
	"alta/book-api/config"
	"alta/book-api/models"
	"alta/book-api/util"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	echo "github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestGetUsers(t *testing.T) {
	// create database connection and create controller
	config := config.GetConfig()
	db := util.MysqlDatabaseConnection(config)
	userModel := models.NewUserModel(db)
	userController := NewController(userModel)

	// setting controller
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	userController.GetUser(context)

	// build struct response
	type Response struct {
		Users []struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		} `json:"users"`
	}

	var response Response
	resBody := res.Body.String()
	json.Unmarshal([]byte(resBody), &response)

	t.Run("GET /users", func(t *testing.T) {
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, 1, len(response.Users))
		assert.Equal(t, response.Users[0].Name, "Name Test A")
		assert.Equal(t, response.Users[0].Email, "test@alterra.id")
	})
}

func TestGetUserOne(t *testing.T) {
	// create database connection and create controller
	config := config.GetConfig()
	db := util.MysqlDatabaseConnection(config)
	userModel := models.NewUserModel(db)
	userController := NewController(userModel)

	// setting controller
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/users/:id")
	context.SetParamNames("id")
	context.SetParamValues("1")

	userController.GetUserOne(context)

	// Unmarshal respose string to struct
	var response response.GetUserOneResponse
	resBody := res.Body.String()
	json.Unmarshal([]byte(resBody), &response)

	t.Run("GET /users/:id", func(t *testing.T) {
		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, "Name Test A", response.Name)
		assert.Equal(t, "test@alterra.id", response.Email)
	})
}

func setup() {
	// create database connection
	config := config.GetConfig()
	db := util.MysqlDatabaseConnection(config)

	// cleaning data before testing
	db.Migrator().DropTable(&models.User{})
	db.AutoMigrate(&models.User{})

	// preparate dummy data
	var newUser models.User
	newUser.Name = "Name Test A"
	newUser.Email = "test@alterra.id"
	newUser.Password = "password123"

	// user dummy data with model
	userModel := models.NewUserModel(db)
	_, err := userModel.Insert(newUser)
	if err != nil {
		fmt.Println(err)
	}
}
