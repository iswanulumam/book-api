package user

import (
	"alta/book-api/config"
	"alta/book-api/models"
	"alta/book-api/util"
	"encoding/json"
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
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	userController.GetUser(context)

	// build struct response
	type Response struct {
		Users []interface{} `json:"users"`
	}
	var response Response
	resBody := res.Body.String()
	json.Unmarshal([]byte(resBody), &response)

	t.Run("GET /users", func(t *testing.T) {
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, 0, len(response.Users))
	})
}

func setup() {
	// create database connection
	config := config.GetConfig()
	db := util.MysqlDatabaseConnection(config)

	// cleaning data before testing
	db.Migrator().DropTable(&models.User{})
	db.AutoMigrate(&models.User{})
}
