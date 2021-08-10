package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDresponse(t *testing.T) {
	t.Run("func NewInternalServerErrorResponse()", func(t *testing.T) {
		InternalServerError := NewInternalServerErrorResponse()
		assert.Equal(t, InternalServerError.Code, 500)
		assert.Equal(t, InternalServerError.Message, "Internal server error")
	})

	t.Run("func NewNotFoundResponse()", func(t *testing.T) {
		NotFound := NewNotFoundResponse()
		assert.Equal(t, NotFound.Code, 404)
		assert.Equal(t, NotFound.Message, "Not found")
	})

	t.Run("func NewBadRequestResponse()", func(t *testing.T) {
		BadRequest := NewBadRequestResponse()
		assert.Equal(t, BadRequest.Code, 400)
		assert.Equal(t, BadRequest.Message, "Bad request")
	})

	t.Run("func NewConflictResponse()", func(t *testing.T) {
		Conflict := NewConflictResponse()
		assert.Equal(t, Conflict.Code, 409)
		assert.Equal(t, Conflict.Message, "Data has been modified")
	})
}
