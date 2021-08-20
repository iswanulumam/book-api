package customer

type PostCustomerRequest struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Gender   string `json:"gender" form:"gender"`
	Password string `json:"password" form:"password"`
}

type EditCustomerRequest struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Gender   string `json:"gender" form:"gender"`
	Password string `json:"password" form:"password"`
}
