package transfer

import (
	"context"

	"github.com/stephanvebrian/e-pocket/pocket-engine/logic/statemachine"
	handlerModel "github.com/stephanvebrian/e-pocket/pocket-engine/model/handler"
)

func (tl *transferLogic) CreateTransfer(ctx context.Context, request handlerModel.CreateTransferRequest) (handlerModel.CreateTransferResponse, error) {
	// Initialize the state machine
	sm := statemachine.New()

	// Register states
	sm.RegisterState(statemachine.State(StateInit), tl.handleInitState)
	sm.RegisterState(statemachine.State(StateCreate), tl.handleCreateState)
	sm.RegisterState(statemachine.State(StateDeduct), tl.handleDeductState)
	sm.RegisterState(statemachine.State(StateAdd), tl.handleAddState)
	sm.RegisterState(statemachine.State(StateUpdateStatus), tl.handleUpdateStatusState)
	sm.RegisterState(statemachine.State(StateTransactionHistory), tl.handleTransactionHistoryState)
	sm.RegisterState(statemachine.State(StateComplete), tl.handleCompleteState)

	// Initialize the state transition
	transition := &TransferStateTransition{
		State:   StateInit,
		Request: request,
	}

	// Run the state machine
	finalTransition, err := sm.Run(ctx, statemachine.State(StateInit), transition)
	if err != nil {
		return handlerModel.CreateTransferResponse{}, err
	}

	// Return the response
	finalTransfer := finalTransition.(*TransferStateTransition).Transfer
	return handlerModel.CreateTransferResponse{
		IdempotencyKey: finalTransfer.ReferenceID,
		TransactionID:  finalTransfer.ID.String(),
		Status:         string(finalTransfer.Status),
	}, nil
}
