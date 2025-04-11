package model

import (
	"fmt"
	"net/http"
)

type ErrorCode string

const (
	// common error code
	ValidationError   ErrorCode = "VALIDATION_ERROR"
	InvalidBody       ErrorCode = "INVALID_BODY"
	InvalidParameter  ErrorCode = "INVALID_PARAMETER"
	UnexpectedError   ErrorCode = "UNEXPECTED_ERROR"
	DatabaseError     ErrorCode = "DATABASE_ERROR"
	DataNotFoundError ErrorCode = "NOT_FOUND_ERROR"
	NotPermittedError ErrorCode = "NOT_PERMITTED"

	// transfer error code
	TransferAlreadyExists ErrorCode = "TRANSFER_ALREADY_EXISTS"
	InsufficientBalance   ErrorCode = "INSUFFICIENT_BALANCE"
)

type ErrorResponse struct {
	HTTPCode  int         `json:"-"`         // HTTP status code
	Timestamp string      `json:"timestamp"` // Timestamp of the error, will be filled on middleware layer
	Code      ErrorCode   `json:"code"`      // Unique error code
	Message   string      `json:"message"`   // Human-readable error message
	Details   interface{} `json:"details"`   // Optional additional details (e.g., validation errors)
}

func (e ErrorResponse) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

type ErrorResponseOption struct {
	Writer     http.ResponseWriter
	Request    *http.Request
	StatusCode int
	Response   ErrorResponse
}
