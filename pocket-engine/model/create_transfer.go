package model

type CreateTransferRequest struct {
	IdempotencyKey string                 `json:"idempotencyKey" validate:"required,len=36"`
	Sender         TransferAccountRequest `json:"sender" validate:"required"`
	Receiver       TransferAccountRequest `json:"receiver" validate:"required"`
	Amount         uint64                 `json:"amount" validate:"required,gt=0"`
}

type TransferAccountRequest struct {
	Number string `json:"number" validate:"required"`
	Name   string `json:"name" validate:"required"`
}

type CreateTransferResponse struct {
	IdempotencyKey string `json:"idempotencyKey"`
	TransactionID  string `json:"transactionID"`
	Status         string `json:"status"`
}
