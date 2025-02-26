package account

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/stephanvebrian/e-pocket/pocket-engine/model"
	handlerModel "github.com/stephanvebrian/e-pocket/pocket-engine/model/handler"
	"gorm.io/gorm"
)

func (al *accountLogic) CreateAccount(ctx context.Context, request handlerModel.CreateAccountRequest) (handlerModel.CreateAccountResponse, error) {
	userID, err := uuid.Parse(request.UserID)
	if err != nil {
		return handlerModel.CreateAccountResponse{}, model.ErrorResponse{
			HTTPCode: http.StatusBadRequest,
			Code:     model.InvalidBody,
			Message:  "Invalid request body",
		}
	}

	// validate UserID
	var user model.User
	result := al.db.First(&user, "id = ?", request.UserID)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return handlerModel.CreateAccountResponse{}, model.ErrorResponse{
			HTTPCode: http.StatusInternalServerError,
			Code:     model.DatabaseError,
			Message:  "Error when querying user",
		}
	}

	// retrieve latest account pocket
	var account model.Account
	result = al.db.First(&account, "user_id = ?", request.UserID).Order("pocket_number DESC")
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return handlerModel.CreateAccountResponse{}, model.ErrorResponse{
			HTTPCode: http.StatusInternalServerError,
			Code:     model.DatabaseError,
			Message:  "Error when querying account",
		}
	}

	prefix := ""
	suffix := ""
	pocketNumber := 1
	if result.RowsAffected > 0 {
		pocketNumber = account.PocketNumber + 1
		prefix = account.Prefix
		suffix = account.Suffix
	} else {
		// generate prefix & suffix
		currentTime := time.Now()

		// Format the time as YYMMDDHHMMSS for the prefix (using last 2 digits of the year)
		prefix = currentTime.Format("060102150405") // "06" gives the last 2 digits of the year

		// Extract the millisecond part of the time (last 4 digits of UnixNano)
		millis := currentTime.UnixNano() / int64(time.Millisecond) % 10000
		suffix = fmt.Sprintf("%04d", millis)
	}

	createAccount := model.Account{
		AccountNumber: fmt.Sprintf("%s%s%02d", prefix, suffix, pocketNumber),
		Prefix:        prefix,
		Suffix:        suffix,
		PocketNumber:  pocketNumber,
		AccountName:   request.Name,
		Balance:       0,
		Status:        model.AccountStatusActive,
		UserID:        userID,
	}

	createResult := al.db.Create(&createAccount)
	if createResult.Error != nil {
		return handlerModel.CreateAccountResponse{}, model.ErrorResponse{
			HTTPCode: http.StatusInternalServerError,
			Code:     model.DatabaseError,
			Message:  "Error when creating account",
		}
	}

	return handlerModel.CreateAccountResponse{
		AccountNumber: createAccount.AccountNumber,
		Name:          createAccount.AccountName,
		Balance:       createAccount.Balance,
		Status:        string(createAccount.Status),
	}, nil
}
