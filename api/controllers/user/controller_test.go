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
	c := e.NewContext(req, res)
	userController.GetUser(c)

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
