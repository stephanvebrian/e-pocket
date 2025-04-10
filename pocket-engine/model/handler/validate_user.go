package handler

type ValidateUserRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type ValidateUserResponse struct {
	IsValid bool   `json:"isValid"`
	UserID  string `json:"userID"`
}
