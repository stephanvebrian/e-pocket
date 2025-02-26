package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type TransferStatus string

const (
	TransferStatusAuthorized TransferStatus = "AUTHORIZED"
	TransferStatusProcessing TransferStatus = "PROCESSING"
	TransferStatusCaptured   TransferStatus = "CAPTURED"
	TransferStatusCompleted  TransferStatus = "COMPLETED"
	TransferStatusFailed     TransferStatus = "FAILED"
)

type Transfer struct {
	ID              uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	ReferenceID     string
	SenderAccount   string
	Sender          *TransferAccount
	ReceiverAccount string
	Receiver        *TransferAccount
	Amount          uint64
	Status          TransferStatus
	UserID          uuid.UUID
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func (Transfer) TableName() string {
	return "transfer"
}

type TransferAccount struct {
	Name string
}

func (ta *TransferAccount) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	result := TransferAccount{}
	err := json.Unmarshal(bytes, &result)
	*ta = TransferAccount(result)
	return err
}

func (ta *TransferAccount) Value() (driver.Value, error) {
	if ta == nil {
		return nil, nil
	}

	return json.Marshal(ta)
}
