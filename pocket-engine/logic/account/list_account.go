package account

import (
	"context"
	"net/http"

	"github.com/stephanvebrian/e-pocket/pocket-engine/model"
	handlerModel "github.com/stephanvebrian/e-pocket/pocket-engine/model/handler"
	"gorm.io/gorm"
)

func (al *accountLogic) ListAccount(ctx context.Context, request handlerModel.ListAccountRequest) (handlerModel.ListAccountResponse, error) {
	var accounts []model.Account
	result := al.db.Find(&accounts, "user_id = ?", request.UserID)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return handlerModel.ListAccountResponse{}, model.ErrorResponse{
			HTTPCode: http.StatusInternalServerError,
			Code:     model.DatabaseError,
			Message:  "Error when querying account",
		}
	}

	accountsResponse := []handlerModel.AccountData{}
	for _, account := range accounts {
		accountsResponse = append(accountsResponse, handlerModel.AccountData{
			AccountNumber: account.AccountNumber,
			AccountName:   account.AccountName,
			Balance:       account.Balance,
			Status:        string(account.Status),
		})
	}

	return handlerModel.ListAccountResponse{
		Accounts: accountsResponse,
	}, nil
}
