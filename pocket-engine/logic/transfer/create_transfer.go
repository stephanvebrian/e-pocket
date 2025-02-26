package transfer

import (
	"context"
	"net/http"

	"github.com/stephanvebrian/e-pocket/pocket-engine/model"
	handlerModel "github.com/stephanvebrian/e-pocket/pocket-engine/model/handler"
	"gorm.io/gorm"
)

// TODO: Add error handling, ensure the transaction can be marked as failed or canceled, and make the deduction and addition of funds atomic.
func (tl *transferLogic) CreateTransfer(ctx context.Context, request handlerModel.CreateTransferRequest) (handlerModel.CreateTransferResponse, error) {
	var transfer model.Transfer
	result := tl.db.First(&transfer).Where("reference_id = ?", request.IdempotencyKey)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return handlerModel.CreateTransferResponse{}, model.ErrorResponse{
			HTTPCode: http.StatusInternalServerError,
			Code:     model.DatabaseError,
			Message:  "Error when querying transfer",
		}
	}

	// transfer exists
	if result.RowsAffected > 0 {
		// flow existing transfer

		// validate user id, if its not their transfer
		if transfer.UserID.String() != request.UserID {
			return handlerModel.CreateTransferResponse{}, model.ErrorResponse{
				HTTPCode: http.StatusForbidden,
				Code:     model.NotPermittedError,
				Message:  "Forbidden",
			}
		}

		// otherwise return the existing model from result
		return handlerModel.CreateTransferResponse{
			IdempotencyKey: transfer.ReferenceID,
			TransactionID:  transfer.ID.String(),
			Status:         string(transfer.Status),
		}, nil
	}

	// 1. initate the transfer
	createTransfer := model.Transfer{
		ReferenceID:   request.IdempotencyKey,
		SenderAccount: request.Sender.Number,
		Sender: &model.TransferAccount{
			Name: request.Sender.Name,
		},
		ReceiverAccount: request.Receiver.Number,
		Receiver: &model.TransferAccount{
			Name: request.Receiver.Name,
		},
		Amount: request.Amount,
		Status: model.TransferStatusProcessing,
	}

	createResult := tl.db.Create(&createTransfer)
	if createResult.Error != nil {
		return handlerModel.CreateTransferResponse{}, model.ErrorResponse{
			HTTPCode: http.StatusInternalServerError,
			Code:     model.DatabaseError,
			Message:  "Error when creating transfer",
		}
	}

	// 2. retreive sender money
	var senderAccount model.Account
	senderResult := tl.db.First(&senderAccount).Where("account_number = ?", request.Sender.Number)
	if senderResult.Error != nil {
		if senderResult.Error == gorm.ErrRecordNotFound {
			return handlerModel.CreateTransferResponse{}, model.ErrorResponse{
				HTTPCode: http.StatusNotFound,
				Code:     model.DataNotFoundError,
				Message:  "Sender account not found",
			}
		}

		return handlerModel.CreateTransferResponse{}, model.ErrorResponse{
			HTTPCode: http.StatusInternalServerError,
			Code:     model.DatabaseError,
			Message:  "Error when querying sender account",
		}
	}

	// 2.1 check if sender has enough money, if not return error
	if senderAccount.Balance < request.Amount {
		return handlerModel.CreateTransferResponse{}, model.ErrorResponse{
			HTTPCode: http.StatusUnprocessableEntity,
			Code:     model.InsufficientBalance,
			Message:  "Insufficient balance",
		}
	}

	// 2.2 subtract the money from sender
	deductResult := tl.db.Model(&model.Account{}).Where("id = ?", senderAccount.ID).Update("balance", gorm.Expr("balance - ?", request.Amount))
	if deductResult.Error != nil {
		return handlerModel.CreateTransferResponse{}, model.ErrorResponse{
			HTTPCode: http.StatusInternalServerError,
			Code:     model.DatabaseError,
			Message:  "Error when deducting customer balance",
		}
	}

	if deductResult.RowsAffected != 1 {
		return handlerModel.CreateTransferResponse{}, model.ErrorResponse{
			HTTPCode: http.StatusInternalServerError,
			Code:     model.DatabaseError,
			Message:  "Error when deducting customer balance",
		}
	}

	// 3. add the money to receiver
	additionResult := tl.db.Model(&model.Account{}).Where("account_number = ?", request.Receiver.Number).Update("balance", gorm.Expr("balance + ?", request.Amount))
	if additionResult.Error != nil {
		return handlerModel.CreateTransferResponse{}, model.ErrorResponse{
			HTTPCode: http.StatusInternalServerError,
			Code:     model.DatabaseError,
			Message:  "Error when adding customer balance",
		}
	}

	if additionResult.RowsAffected != 1 {
		return handlerModel.CreateTransferResponse{}, model.ErrorResponse{
			HTTPCode: http.StatusInternalServerError,
			Code:     model.DatabaseError,
			Message:  "Error when adding customer balance",
		}
	}

	// 4. update the transfer status to success
	updateStatusResult := tl.db.Model(&model.Transfer{}).Where("reference_id = ?", request.IdempotencyKey).Update("status", model.TransferStatusCompleted)
	if updateStatusResult.Error != nil {
		return handlerModel.CreateTransferResponse{}, model.ErrorResponse{
			HTTPCode: http.StatusInternalServerError,
			Code:     model.DatabaseError,
			Message:  "Error when updating transfer status",
		}
	}

	if updateStatusResult.RowsAffected != 1 {
		return handlerModel.CreateTransferResponse{}, model.ErrorResponse{
			HTTPCode: http.StatusInternalServerError,
			Code:     model.DatabaseError,
			Message:  "Error when updating transfer status",
		}
	}

	// 5. return the response
	var updatedTransfer model.Transfer
	transferResult := tl.db.First(&updatedTransfer).Where("reference_id = ?", request.IdempotencyKey)
	if transferResult.Error != nil {
		return handlerModel.CreateTransferResponse{}, model.ErrorResponse{
			HTTPCode: http.StatusInternalServerError,
			Code:     model.DatabaseError,
			Message:  "Error when querying transfer",
		}
	}

	return handlerModel.CreateTransferResponse{
		IdempotencyKey: updatedTransfer.ReferenceID,
		TransactionID:  updatedTransfer.ID.String(),
		Status:         string(updatedTransfer.Status),
	}, nil
}
