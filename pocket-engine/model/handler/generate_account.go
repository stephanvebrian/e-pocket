package handler

type GenerateAccountRequest struct {
}

type GenerateAccountResponse struct {
	UserID        string `json:"userId"`
	Username      string `json:"username"`
	Password      string `json:"password"`
	AccountNumber string `json:"accountNumber"`
	Name          string `json:"name"`
	Balance       uint64 `json:"balance"`
	Status        string `json:"status"`
}
