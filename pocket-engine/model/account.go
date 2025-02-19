package model

import "time"

type AccountStatus string

const (
	AccountStatusActive   AccountStatus = "ACTIVE"
	AccountStatusInactive AccountStatus = "INACTIVE"
)

type Account struct {
	ID            uint64 `gorm:"primarykey"`
	AccountNumber string
	AccountName   string
	Balance       uint64
	Status        AccountStatus
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (Account) TableName() string {
	return "account"
}
