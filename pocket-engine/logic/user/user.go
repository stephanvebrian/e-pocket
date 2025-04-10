package user

import (
	"context"

	handlerModel "github.com/stephanvebrian/e-pocket/pocket-engine/model/handler"
	"gorm.io/gorm"
)

type userLogic struct {
	db *gorm.DB
}

type UserLogic interface {
	ValidateUser(ctx context.Context, request handlerModel.ValidateUserRequest) (handlerModel.ValidateUserResponse, error)
}

type UserLogicOptions struct {
	DB *gorm.DB
}

func New(opts UserLogicOptions) UserLogic {
	return &userLogic{
		db: opts.DB,
	}
}
