package response

import (
	"alta/book-api/models"
)

type GetUserResponse struct {
	Users []models.User `json:"users"`
}

func NewGetUserResponse(users []models.User) GetUserResponse {
	return GetUserResponse{
		Users: users,
	}
}

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
