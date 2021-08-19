package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDresponse(t *testing.T) {
	t.Run("func NewSuccessOperationResponse()", func(t *testing.T) {
		SuccessOperation := NewSuccessOperationResponse()
		assert.Equal(t, SuccessOperation.Message, "success operation")
	})

	t.Run("func NewInternalServerErrorResponse()", func(t *testing.T) {
		InternalServerError := NewInternalServerErrorResponse()
		assert.Equal(t, InternalServerError.Message, "internal server error")
	})

	t.Run("func NewNotFoundResponse()", func(t *testing.T) {
		NotFound := NewNotFoundResponse()
		assert.Equal(t, NotFound.Message, "not found")
	})

	t.Run("func NewBadRequestResponse()", func(t *testing.T) {
		BadRequest := NewBadRequestResponse()
		assert.Equal(t, BadRequest.Message, "bad request")
	})

	t.Run("func NewConflictResponse()", func(t *testing.T) {
		Conflict := NewConflictResponse()
		assert.Equal(t, Conflict.Message, "data has been modified")
	})
}
