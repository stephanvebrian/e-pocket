package user

import (
	"context"
	"net/http"

	"github.com/stephanvebrian/e-pocket/pocket-engine/model"
	handlerModel "github.com/stephanvebrian/e-pocket/pocket-engine/model/handler"
	"gorm.io/gorm"
)

func (ul *userLogic) ValidateUser(ctx context.Context, request handlerModel.ValidateUserRequest) (handlerModel.ValidateUserResponse, error) {
	var user model.User
	result := ul.db.First(&user, "username = ?", request.Username)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return handlerModel.ValidateUserResponse{}, model.ErrorResponse{
			HTTPCode: http.StatusInternalServerError,
			Code:     model.DatabaseError,
			Message:  "Error when querying user",
		}
	}

	if result.Error == gorm.ErrRecordNotFound {
		return handlerModel.ValidateUserResponse{}, model.ErrorResponse{
			HTTPCode: http.StatusNotFound,
			Code:     model.DataNotFoundError,
			Message:  "User not found",
		}
	}

	if user.Password != request.Password {
		return handlerModel.ValidateUserResponse{}, model.ErrorResponse{
			HTTPCode: http.StatusUnauthorized,
			Code:     model.NotPermittedError,
			Message:  "Invalid credentials",
		}
	}

	return handlerModel.ValidateUserResponse{
		IsValid: true,
		UserID:  user.ID.String(),
	}, nil
}
