package common

//DefaultResponse default payload response
type DefaultResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

//NewInternalServerErrorResponse default internal server error response
func NewSuccessOperationResponse() DefaultResponse {
	return DefaultResponse{
		200,
		"Successful Operation",
	}
}

//NewInternalServerErrorResponse default internal server error response
func NewInternalServerErrorResponse() DefaultResponse {
	return DefaultResponse{
		500,
		"Internal Server Error",
	}
}

//NewNotFoundResponse default not found error response
func NewNotFoundResponse() DefaultResponse {
	return DefaultResponse{
		404,
		"Not Found",
	}
}

//NewBadRequestResponse default not found error response
func NewBadRequestResponse() DefaultResponse {
	return DefaultResponse{
		400,
		"Bad Request",
	}
}

//NewConflictResponse default not found error response
func NewConflictResponse() DefaultResponse {
	return DefaultResponse{
		409,
		"Data Has Been Modified",
	}
}
