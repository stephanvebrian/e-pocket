package handler

type ListAccountRequest struct {
	UserID string `json:"userID"`
}

type ListAccountResponse struct {
	Accounts []AccountData `json:"accounts"`
}

type AccountData struct {
	AccountNumber string `json:"accountNumber"`
	AccountName   string `json:"accountName"`
	Balance       uint64 `json:"balance"`
	Status        string `json:"status"`
}
