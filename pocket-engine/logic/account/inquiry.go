package account

import (
	"context"
	"net/http"

	"github.com/stephanvebrian/e-pocket/pocket-engine/model"
	handlerModel "github.com/stephanvebrian/e-pocket/pocket-engine/model/handler"
	"gorm.io/gorm"
)

func (al *accountLogic) Inquiry(ctx context.Context, request handlerModel.InquiryAccountRequest) (handlerModel.InquiryAccountResponse, error) {
	var account model.Account
	result := al.db.First(&account, "account_number = ? AND status = ?", request.AccountNumber, model.AccountStatusActive)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return handlerModel.InquiryAccountResponse{}, model.ErrorResponse{
			HTTPCode: http.StatusInternalServerError,
			Code:     model.DatabaseError,
			Message:  "Error when querying account",
		}
	}

	if result.Error == gorm.ErrRecordNotFound {
		return handlerModel.InquiryAccountResponse{}, model.ErrorResponse{
			HTTPCode: http.StatusNotFound,
			Code:     model.DataNotFoundError,
			Message:  "Account not found",
		}
	}

	return handlerModel.InquiryAccountResponse{
		AccountNumber: account.AccountNumber,
		AccountName:   account.AccountName,
	}, nil
}
