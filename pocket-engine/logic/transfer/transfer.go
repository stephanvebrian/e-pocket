package transfer

import (
	"context"

	handlerModel "github.com/stephanvebrian/e-pocket/pocket-engine/model/handler"
	"gorm.io/gorm"
)

type transferLogic struct {
	db *gorm.DB
}

type TransferLogic interface {
	CreateTransfer(ctx context.Context, request handlerModel.CreateTransferRequest) (handlerModel.CreateTransferResponse, error)
}

type TransferLogicOptions struct {
	DB *gorm.DB
}

func New(opts TransferLogicOptions) TransferLogic {
	return &transferLogic{
		db: opts.DB,
	}
}
