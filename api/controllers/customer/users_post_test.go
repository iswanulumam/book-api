package customer

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
	userModel := models.NewCustomerModel(db)
	customerController := NewController(userModel)

	// input controller
	reqBody, _ := json.Marshal(map[string]string{
		"name":     "Name Test",
		"email":    "test@alterra.id",
		"password": "test123",
	})

	// setting controller
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
	res := httptest.NewRecorder()
	req.Header.Set("Content-Type", "application/json")
	context := e.NewContext(req, res)
	customerController.PostOne(context)

	// build struct response
	type Response struct {
		Message string `json:"message"`
	}
	var response Response
	resBody := res.Body.String()
	json.Unmarshal([]byte(resBody), &response)

	// testing stuff
	t.Run("POST /customers", func(t *testing.T) {
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, "success operation", response.Message)
	})
}
