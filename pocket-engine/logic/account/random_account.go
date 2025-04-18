package account

import (
	"context"
	"net/http"

	"github.com/stephanvebrian/e-pocket/pocket-engine/model"
	handlerModel "github.com/stephanvebrian/e-pocket/pocket-engine/model/handler"
	"gorm.io/gorm"
)

func (al *accountLogic) RandomAccount(ctx context.Context, request handlerModel.RandomAccountRequest) (handlerModel.RandomAccountResponse, error) {
	var account model.Account
	err := al.db.Where("user_id != ?", request.UserID).Order("RANDOM()").First(&account).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return handlerModel.RandomAccountResponse{}, model.ErrorResponse{
				HTTPCode: http.StatusNotFound,
				Code:     model.DataNotFoundError,
				Message:  "Account not found",
			}
		}
		return handlerModel.RandomAccountResponse{}, model.ErrorResponse{
			HTTPCode: http.StatusInternalServerError,
			Code:     model.DatabaseError,
			Message:  "Error when querying account",
		}
	}

	accountResponse := handlerModel.AccountData{
		AccountNumber: account.AccountNumber,
		AccountName:   account.AccountName,
		Balance:       account.Balance,
		Status:        string(account.Status),
	}

	return handlerModel.RandomAccountResponse{
		Account: accountResponse,
	}, nil
}
