package user

import (
	"alta/book-api/config"
	"alta/book-api/models"
	"alta/book-api/util"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	echo "github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestPostUsers(t *testing.T) {
	// create database connection and create controller
	config := config.GetConfig()
	db := util.MysqlDatabaseConnection(config)
	userModel := models.NewUserModel(db)
	userController := NewController(userModel)

	// input controller
	reqBody, _ := json.Marshal(map[string]string{
		"name":     "Iswanul Umam",
		"email":    "umam@alterra.id",
		"password": "admin123",
	})

	// setting controller
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(reqBody))
	res := httptest.NewRecorder()
	req.Header.Set("Content-Type", "application/json")
	context := e.NewContext(req, res)
	userController.PostUser(context)

	// build struct response
	type Response struct {
		Users []interface{} `json:"users"`
	}
	var response Response
	resBody := res.Body.String()
	json.Unmarshal([]byte(resBody), &response)

	// testing stuff
	t.Run("POST /users", func(t *testing.T) {
		assert.Equal(t, 200, res.Code)
	})
}
