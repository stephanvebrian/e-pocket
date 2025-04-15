package model

import (
	"time"

	"github.com/google/uuid"
)

type TransactionHistoryStatus string

const (
	TransactionHistoryStatusPending   TransactionHistoryStatus = "PENDING"
	TransactionHistoryStatusCompleted TransactionHistoryStatus = "COMPLETED"
	TransactionHistoryStatusFailed    TransactionHistoryStatus = "FAILED"
)

type TransactionHistoryType string

const (
	TransactionHistoryTypeOutgoing TransactionHistoryType = "OUTGOING"
	TransactionHistoryTypeIncoming TransactionHistoryType = "INCOMING"
)

type TransactionHistory struct {
	ID                uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	UserID            uuid.UUID
	AccountID         uint64
	TransactionType   TransactionHistoryType
	TransactionAmount uint64
	EndingBalance     uint64
	Status            TransactionHistoryStatus
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

func (TransactionHistory) TableName() string {
	return "transaction_history"
}
