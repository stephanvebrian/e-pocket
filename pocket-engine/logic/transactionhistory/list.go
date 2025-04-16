package transactionhistory

import (
	"context"
	"net/http"

	"github.com/stephanvebrian/e-pocket/pocket-engine/model"
	handlerModel "github.com/stephanvebrian/e-pocket/pocket-engine/model/handler"
	"gorm.io/gorm"
)

func (th *transactionHistoryLogic) ListTransactionHistory(ctx context.Context, request handlerModel.ListTransactionHistoryRequest) (handlerModel.ListTransactionHistoryResponse, error) {
	var transactionHistory []model.TransactionHistory
	result := th.db.Find(&transactionHistory, "user_id = ?", request.UserID)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return handlerModel.ListTransactionHistoryResponse{}, model.ErrorResponse{
			HTTPCode: http.StatusInternalServerError,
			Code:     model.DatabaseError,
			Message:  "Error when querying transaction history",
		}
	}

	transactionHistoryResponse := []handlerModel.TransactionHistoryData{}
	for _, transaction := range transactionHistory {
		transactionHistoryResponse = append(transactionHistoryResponse, handlerModel.TransactionHistoryData{
			ID:     transaction.ID.String(),
			UserID: transaction.UserID.String(),
			Account: handlerModel.AccountData{
				AccountNumber: "",
			},
			TransactionType: string(transaction.TransactionType),
			Amount:          transaction.TransactionAmount,
			EndingBalance:   transaction.EndingBalance,
			Status:          string(transaction.Status),
			CreatedAt:       transaction.CreatedAt.String(),
			UpdatedAt:       transaction.UpdatedAt.String(),
		})
	}

	return handlerModel.ListTransactionHistoryResponse{
		TransactionHistory: transactionHistoryResponse,
	}, nil
}
