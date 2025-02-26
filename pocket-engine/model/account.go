package model

import (
	"time"

	"github.com/google/uuid"
)

type AccountStatus string

const (
	AccountStatusActive   AccountStatus = "ACTIVE"
	AccountStatusInactive AccountStatus = "INACTIVE"
)

type Account struct {
	ID            uint64 `gorm:"primarykey"`
	AccountNumber string
	Prefix        string
	Suffix        string
	PocketNumber  int
	AccountName   string
	Balance       uint64
	Status        AccountStatus
	UserID        uuid.UUID
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (Account) TableName() string {
	return "account"
}
