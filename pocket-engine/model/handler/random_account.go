package handler

type RandomAccountRequest struct {
	UserID string `json:"userID"`
}

type RandomAccountResponse struct {
	Account AccountData `json:"account"`
}
