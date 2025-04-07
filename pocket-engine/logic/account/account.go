package account

import (
	"context"

	handlerModel "github.com/stephanvebrian/e-pocket/pocket-engine/model/handler"
	"gorm.io/gorm"
)

type accountLogic struct {
	db *gorm.DB
}

type AccountLogic interface {
	GenerateAccount(ctx context.Context, request handlerModel.GenerateAccountRequest) (handlerModel.GenerateAccountResponse, error)
}

type AccountLogicOptions struct {
	DB *gorm.DB
}

func New(opts AccountLogicOptions) AccountLogic {
	return &accountLogic{
		db: opts.DB,
	}
}
