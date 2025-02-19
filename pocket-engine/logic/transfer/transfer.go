package transfer

import (
	"context"

	"github.com/stephanvebrian/e-pocket/pocket-engine/model"
	"gorm.io/gorm"
)

type transferLogic struct {
	db *gorm.DB
}

type TransferLogic interface {
	CreateTransfer(ctx context.Context, request model.CreateTransferRequest) (model.CreateTransferResponse, error)
}

type TransferLogicOptions struct {
	DB *gorm.DB
}

func New(opts TransferLogicOptions) TransferLogic {
	return &transferLogic{
		db: opts.DB,
	}
}
