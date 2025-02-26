package handler

type CreateAccountRequest struct {
	Name   string `json:"name" validate:"required"`
	UserID string `json:"userID" validate:"required"` // TODO: temporary user id, move it to headers when auth is implemented
}

type CreateAccountResponse struct {
	AccountNumber string `json:"accountNumber"`
	Name          string `json:"name"`
	Balance       uint64 `json:"balance"`
	Status        string `json:"status"`
}
