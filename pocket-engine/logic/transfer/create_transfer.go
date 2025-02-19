package transfer

import (
	"context"
	"strings"

	"github.com/google/uuid"
	"github.com/stephanvebrian/e-pocket/pocket-engine/model"
	"gorm.io/gorm"
)

// TODO: Add error handling, ensure the transaction can be marked as failed or canceled, and make the deduction and addition of funds atomic.
func (tl *transferLogic) CreateTransfer(ctx context.Context, request model.CreateTransferRequest) (model.CreateTransferResponse, error) {
	var transfer model.Transfer
	result := tl.db.First(&transfer).Where("reference_id = ?", request.IdempotencyKey)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return model.CreateTransferResponse{}, model.ErrorResponse{
			Code:    model.DatabaseError,
			Message: "Error when querying transfer",
		}
	}

	// transfer exists
	if result.RowsAffected > 0 {
		// flow existing transfer

		// validate user id, if its not their transfer, return error
		// otherwise return the existing model from result
		return model.CreateTransferResponse{
			IdempotencyKey: transfer.ReferenceID,
			TransactionID:  transfer.TransactionID,
			Status:         string(transfer.Status),
		}, nil
	}

	transactionID := strings.Replace(uuid.NewString(), "-", "", -1)

	// 1. initate the transfer
	createTransfer := model.Transfer{
		ReferenceID:   request.IdempotencyKey,
		TransactionID: transactionID,
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
		return model.CreateTransferResponse{}, model.ErrorResponse{
			Code:    model.DatabaseError,
			Message: "Error when creating transfer",
		}
	}

	// 2. retreive sender money
	var senderAccount model.Account
	senderResult := tl.db.First(&senderAccount).Where("account_number = ?", request.Sender.Number)
	if senderResult.Error != nil {
		if senderResult.Error == gorm.ErrRecordNotFound {
			return model.CreateTransferResponse{}, model.ErrorResponse{
				Code:    model.DataNotFoundError,
				Message: "Sender account not found",
			}
		}

		return model.CreateTransferResponse{}, model.ErrorResponse{
			Code:    model.DatabaseError,
			Message: "Error when querying sender account",
		}
	}

	// 2.1 check if sender has enough money, if not return error
	if senderAccount.Balance < request.Amount {
		return model.CreateTransferResponse{}, model.ErrorResponse{
			Code:    model.InsufficientBalance,
			Message: "Insufficient balance",
		}
	}

	// 2.2 subtract the money from sender
	deductResult := tl.db.Model(&model.Account{}).Where("id = ?", senderAccount.ID).Update("balance", gorm.Expr("balance - ?", request.Amount))
	if deductResult.Error != nil {
		return model.CreateTransferResponse{}, model.ErrorResponse{
			Code:    model.DatabaseError,
			Message: "Error when deducting customer balance",
		}
	}

	if deductResult.RowsAffected != 1 {
		return model.CreateTransferResponse{}, model.ErrorResponse{
			Code:    model.DatabaseError,
			Message: "Error when deducting customer balance",
		}
	}

	// 3. add the money to receiver
	additionResult := tl.db.Model(&model.Account{}).Where("account_number = ?", request.Receiver.Number).Update("balance", gorm.Expr("balance + ?", request.Amount))
	if additionResult.Error != nil {
		return model.CreateTransferResponse{}, model.ErrorResponse{
			Code:    model.DatabaseError,
			Message: "Error when adding customer balance",
		}
	}

	if additionResult.RowsAffected != 1 {
		return model.CreateTransferResponse{}, model.ErrorResponse{
			Code:    model.DatabaseError,
			Message: "Error when adding customer balance",
		}
	}

	// 4. update the transfer status to success
	updateStatusResult := tl.db.Model(&model.Transfer{}).Where("reference_id = ?", request.IdempotencyKey).Update("status", model.TransferStatusCompleted)
	if updateStatusResult.Error != nil {
		return model.CreateTransferResponse{}, model.ErrorResponse{
			Code:    model.DatabaseError,
			Message: "Error when updating transfer status",
		}
	}

	if updateStatusResult.RowsAffected != 1 {
		return model.CreateTransferResponse{}, model.ErrorResponse{
			Code:    model.DatabaseError,
			Message: "Error when updating transfer status",
		}
	}

	// 5. return the response
	var updatedTransfer model.Transfer
	transferResult := tl.db.First(&updatedTransfer).Where("reference_id = ?", request.IdempotencyKey)
	if transferResult.Error != nil {
		return model.CreateTransferResponse{}, model.ErrorResponse{
			Code:    model.DatabaseError,
			Message: "Error when querying transfer",
		}
	}

	return model.CreateTransferResponse{
		IdempotencyKey: updatedTransfer.ReferenceID,
		TransactionID:  updatedTransfer.TransactionID,
		Status:         string(updatedTransfer.Status),
	}, nil
}
