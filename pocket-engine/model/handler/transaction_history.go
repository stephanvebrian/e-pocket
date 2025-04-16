package handler

type ListTransactionHistoryRequest struct {
	UserID string `json:"userID"`
}

type ListTransactionHistoryResponse struct {
	TransactionHistory []TransactionHistoryData `json:"transactionHistory"`
}

type TransactionHistoryData struct {
	ID              string      `json:"id"`
	UserID          string      `json:"userID"`
	Account         AccountData `json:"account"`
	TransactionType string      `json:"transactionType"`
	Amount          uint64      `json:"amount"`
	EndingBalance   uint64      `json:"endingBalance"`
	Status          string      `json:"status"`
	CreatedAt       string      `json:"createdAt"`
	UpdatedAt       string      `json:"updatedAt"`
}
