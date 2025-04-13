package transfer

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/stephanvebrian/e-pocket/pocket-engine/logic/statemachine"
	"github.com/stephanvebrian/e-pocket/pocket-engine/model"
	handlerModel "github.com/stephanvebrian/e-pocket/pocket-engine/model/handler"
	"gorm.io/gorm"
)

type TransferState = statemachine.State

const (
	StateInit         TransferState = "INIT"
	StateCreate       TransferState = "CREATE"
	StateDeduct       TransferState = "DEDUCT"
	StateAdd          TransferState = "ADD"
	StateUpdateStatus TransferState = "UPDATE_STATUS"
	StateComplete     TransferState = "COMPLETE"
	StateFailed       TransferState = "FAILED"
)

type TransferStateTransition struct {
	State           TransferState
	Transfer        model.Transfer
	Request         handlerModel.CreateTransferRequest
	SenderAccount   model.Account
	ReceiverAccount model.Account
}

func (t *TransferStateTransition) GetState() statemachine.State {
	return statemachine.State(t.State)
}

func (t *TransferStateTransition) SetState(state statemachine.State) {
	t.State = TransferState(state)
}

func (tl *transferLogic) handleInitState(ctx context.Context, args statemachine.StateTransition) (statemachine.StateTransition, error) {
	transition := args.(*TransferStateTransition)

	// Check if transfer already exists
	var transfer model.Transfer
	result := tl.db.First(&transfer, "reference_id = ?", transition.Request.IdempotencyKey)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		fmt.Println("Error when querying transfer:", result.Error)
		return nil, model.ErrorResponse{
			HTTPCode: http.StatusInternalServerError,
			Code:     model.DatabaseError,
			Message:  "Error when querying transfer",
		}
	}

	// If transfer exists, validate and return
	if result.RowsAffected > 0 {
		if transfer.UserID.String() != transition.Request.UserID {
			return nil, model.ErrorResponse{
				HTTPCode: http.StatusForbidden,
				Code:     model.NotPermittedError,
				Message:  "Forbidden",
			}
		}
		transition.Transfer = transfer
		transition.SetState(StateComplete)
		return transition, nil
	}

	// Transition to CREATE state
	transition.SetState(StateCreate)
	return transition, nil
}

func (tl *transferLogic) handleCreateState(ctx context.Context, args statemachine.StateTransition) (statemachine.StateTransition, error) {
	transition := args.(*TransferStateTransition)

	userID, _ := uuid.Parse(transition.Request.UserID)

	// Create the transfer record
	transfer := model.Transfer{
		ReferenceID:   transition.Request.IdempotencyKey,
		SenderAccount: transition.Request.Sender.Number,
		Sender: &model.TransferAccount{
			Name: "",
		},
		ReceiverAccount: transition.Request.Receiver.Number,
		Receiver: &model.TransferAccount{
			Name: "",
		},
		Amount: transition.Request.Amount,
		Status: model.TransferStatusProcessing,
		UserID: userID,
	}

	if err := tl.db.Create(&transfer).Error; err != nil {
		return nil, model.ErrorResponse{
			HTTPCode: http.StatusInternalServerError,
			Code:     model.DatabaseError,
			Message:  "Error when creating transfer",
		}
	}

	transition.Transfer = transfer
	transition.SetState(StateDeduct)
	return transition, nil
}

func (tl *transferLogic) handleDeductState(ctx context.Context, args statemachine.StateTransition) (statemachine.StateTransition, error) {
	transition := args.(*TransferStateTransition)

	// Retrieve sender account
	var senderAccount model.Account
	if err := tl.db.First(&senderAccount, "account_number = ?", transition.Request.Sender.Number).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, model.ErrorResponse{
				HTTPCode: http.StatusNotFound,
				Code:     model.DataNotFoundError,
				Message:  "Sender account not found",
			}
		}
		return nil, model.ErrorResponse{
			HTTPCode: http.StatusInternalServerError,
			Code:     model.DatabaseError,
			Message:  "Error when querying sender account",
		}
	}

	// Check if sender has enough balance
	if senderAccount.Balance < transition.Request.Amount {
		return nil, model.ErrorResponse{
			HTTPCode: http.StatusUnprocessableEntity,
			Code:     model.InsufficientBalance,
			Message:  "Insufficient balance",
		}
	}

	// Deduct from sender
	if err := tl.db.Model(&model.Account{}).Where("id = ?", senderAccount.ID).Update("balance", gorm.Expr("balance - ?", transition.Request.Amount)).Error; err != nil {
		return nil, model.ErrorResponse{
			HTTPCode: http.StatusInternalServerError,
			Code:     model.DatabaseError,
			Message:  "Error when deducting customer balance",
		}
	}

	transition.SenderAccount = senderAccount
	transition.SetState(StateAdd)
	return transition, nil
}

func (tl *transferLogic) handleAddState(ctx context.Context, args statemachine.StateTransition) (statemachine.StateTransition, error) {
	transition := args.(*TransferStateTransition)

	// Add to receiver
	if err := tl.db.Model(&model.Account{}).Where("account_number = ?", transition.Request.Receiver.Number).Update("balance", gorm.Expr("balance + ?", transition.Request.Amount)).Error; err != nil {
		return nil, model.ErrorResponse{
			HTTPCode: http.StatusInternalServerError,
			Code:     model.DatabaseError,
			Message:  "Error when adding customer balance",
		}
	}

	transition.SetState(StateUpdateStatus)
	return transition, nil
}

func (tl *transferLogic) handleUpdateStatusState(ctx context.Context, args statemachine.StateTransition) (statemachine.StateTransition, error) {
	transition := args.(*TransferStateTransition)

	// Update transfer status to completed
	if err := tl.db.Model(&model.Transfer{}).Where("reference_id = ?", transition.Request.IdempotencyKey).Update("status", model.TransferStatusCompleted).Error; err != nil {
		return nil, model.ErrorResponse{
			HTTPCode: http.StatusInternalServerError,
			Code:     model.DatabaseError,
			Message:  "Error when updating transfer status",
		}
	}

	transition.SetState(StateComplete)
	return transition, nil
}

func (tl *transferLogic) handleCompleteState(ctx context.Context, args statemachine.StateTransition) (statemachine.StateTransition, error) {
	transition := args.(*TransferStateTransition)

	// Retrieve the updated transfer
	var updatedTransfer model.Transfer
	if err := tl.db.First(&updatedTransfer, "reference_id = ?", transition.Request.IdempotencyKey).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, model.ErrorResponse{
			HTTPCode: http.StatusInternalServerError,
			Code:     model.DatabaseError,
			Message:  "Error when querying transfer",
			Details:  err,
		}
	}

	transition.Transfer = updatedTransfer
	// TODO: fill transaction history
	return transition, nil
}
