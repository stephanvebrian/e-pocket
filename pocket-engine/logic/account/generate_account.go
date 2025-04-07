package account

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/stephanvebrian/e-pocket/pocket-engine/model"
	handlerModel "github.com/stephanvebrian/e-pocket/pocket-engine/model/handler"
	"gorm.io/gorm"
)

func (al *accountLogic) GenerateAccount(ctx context.Context, request handlerModel.GenerateAccountRequest) (handlerModel.GenerateAccountResponse, error) {
	username := faker.Username()

	// validate UserID
	var user model.User
	result := al.db.First(&user, "username = ?", username)
	if result.Error != gorm.ErrRecordNotFound {
		return handlerModel.GenerateAccountResponse{}, model.ErrorResponse{
			HTTPCode: http.StatusInternalServerError,
			Code:     model.DatabaseError,
			Message:  "Error when querying user",
		}
	}

	if result.RowsAffected > 0 {
		// user already exists, return 500 to be retried
		return handlerModel.GenerateAccountResponse{}, model.ErrorResponse{
			HTTPCode: http.StatusInternalServerError,
			Code:     model.DatabaseError,
			Message:  "User already exists",
		}
	}

	// create an user
	user = model.User{
		Username: username,
		Password: faker.Password(),
	}

	createUserResult := al.db.Create(&user)
	if createUserResult.Error != nil {
		return handlerModel.GenerateAccountResponse{}, model.ErrorResponse{
			HTTPCode: http.StatusInternalServerError,
			Code:     model.DatabaseError,
			Message:  "Error when creating user",
		}
	}

	// create an account
	//
	currentTime := time.Now()

	// Format the time as YYMMDDHHMMSS for the prefix (using last 2 digits of the year)
	prefix := currentTime.Format("060102150405") // "06" gives the last 2 digits of the year
	// Extract the millisecond part of the time (last 4 digits of UnixNano)
	millis := currentTime.UnixNano() / int64(time.Millisecond) % 10000
	suffix := fmt.Sprintf("%04d", millis)
	pocketNumber := 1
	pocketName := "Main Pocket"

	createAccount := model.Account{
		AccountNumber: fmt.Sprintf("%s%s%02d", prefix, suffix, pocketNumber),
		Prefix:        prefix,
		Suffix:        suffix,
		PocketNumber:  pocketNumber,
		AccountName:   pocketName,
		Balance:       0,
		Status:        model.AccountStatusActive,
		UserID:        user.ID,
	}

	createAccountResult := al.db.Create(&createAccount)
	if createAccountResult.Error != nil {
		return handlerModel.GenerateAccountResponse{}, model.ErrorResponse{
			HTTPCode: http.StatusInternalServerError,
			Code:     model.DatabaseError,
			Message:  "Error when creating account",
		}
	}

	return handlerModel.GenerateAccountResponse{
		UserID:        user.ID.String(),
		Username:      user.Username,
		Password:      user.Password,
		AccountNumber: createAccount.AccountNumber,
		Name:          createAccount.AccountName,
		Balance:       createAccount.Balance,
		Status:        string(createAccount.Status),
	}, nil
}
