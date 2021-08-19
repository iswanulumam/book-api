package common

//DefaultResponse default payload response
type DefaultResponse struct {
	Message string `json:"message"`
}

//NewInternalServerErrorResponse default internal server error response
func NewSuccessOperationResponse() DefaultResponse {
	return DefaultResponse{
		"success operation",
	}
}

//NewInternalServerErrorResponse default internal server error response
func NewInternalServerErrorResponse() DefaultResponse {
	return DefaultResponse{
		"internal server error",
	}
}

//NewNotFoundResponse default not found error response
func NewNotFoundResponse() DefaultResponse {
	return DefaultResponse{
		"not found",
	}
}

//NewBadRequestResponse default not found error response
func NewBadRequestResponse() DefaultResponse {
	return DefaultResponse{
		"bad request",
	}
}

//NewConflictResponse default not found error response
func NewConflictResponse() DefaultResponse {
	return DefaultResponse{
		"data has been modified",
	}
}
