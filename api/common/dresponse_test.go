package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDresponse(t *testing.T) {
	t.Run("func NewSuccessOperationResponse()", func(t *testing.T) {
		SuccessOperation := NewSuccessOperationResponse()
		assert.Equal(t, SuccessOperation.Message, "Successful Operation")
	})

	t.Run("func NewInternalServerErrorResponse()", func(t *testing.T) {
		InternalServerError := NewInternalServerErrorResponse()
		assert.Equal(t, InternalServerError.Message, "Internal Server Error")
	})

	t.Run("func NewNotFoundResponse()", func(t *testing.T) {
		NotFound := NewNotFoundResponse()
		assert.Equal(t, NotFound.Message, "Not Found")
	})

	t.Run("func NewBadRequestResponse()", func(t *testing.T) {
		BadRequest := NewBadRequestResponse()
		assert.Equal(t, BadRequest.Message, "Bad Request")
	})

	t.Run("func NewConflictResponse()", func(t *testing.T) {
		Conflict := NewConflictResponse()
		assert.Equal(t, Conflict.Message, "Data Has Been Modified")
	})
}
