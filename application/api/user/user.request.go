package userweb

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	UserType string `json:"type"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
