package response

import (
	"alta/book-api/models"
)

// get users all reponse
type GetUserResponse struct {
	Users []models.User `json:"users"`
}

func NewGetUserResponse(users []models.User) GetUserResponse {
	return GetUserResponse{
		Users: users,
	}
}

// get user one reponse
type GetUserOneResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewGetUserOneResponse(user models.User) GetUserOneResponse {
	return GetUserOneResponse{
		Name:  user.Name,
		Email: user.Email,
	}
}

// post user reponse
type PostUserResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewPostUserResponse(user models.User) PostUserResponse {
	return PostUserResponse{
		Name:  user.Name,
		Email: user.Email,
	}
}
