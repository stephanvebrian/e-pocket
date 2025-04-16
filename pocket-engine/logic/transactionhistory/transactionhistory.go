package transactionhistory

import (
	"context"

	handlerModel "github.com/stephanvebrian/e-pocket/pocket-engine/model/handler"
	"gorm.io/gorm"
)

type transactionHistoryLogic struct {
	db *gorm.DB
}

type TransactionHistoryLogic interface {
	ListTransactionHistory(ctx context.Context, request handlerModel.ListTransactionHistoryRequest) (handlerModel.ListTransactionHistoryResponse, error)
}

type TransactionHistoryLogicOptions struct {
	DB *gorm.DB
}

func New(opts TransactionHistoryLogicOptions) TransactionHistoryLogic {
	return &transactionHistoryLogic{
		db: opts.DB,
	}
}
