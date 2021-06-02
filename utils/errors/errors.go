package errors

import (
	"fmt"
	"net/http"
)

const (
	BadRequest = "bad_request"
	NotFound   = "not_found"
	Forbidden  = "forbidden"
)

type Error struct {
	Type       string
	Message    string
	StatusCode int
	Underlying error
	RequestId  string
}

func (e *Error) Error() string {
	return fmt.Sprintf("%d %s: %s", e.StatusCode, e.Type, e.Message)
}

func NewError(t string, message string, status int, underlying error, requestId string) *Error {
	return &Error{
		Type:       t,
		Message:    message,
		StatusCode: status,
		Underlying: underlying,
		RequestId:  requestId,
	}
}

func NewBadRequestError(message string, underlying error, requestId string) *Error {
	return NewError(BadRequest, message, http.StatusBadRequest, underlying, requestId)
}

func NewNotFoundError(message string, underlying error, requestId string) *Error {
	return NewError(NotFound, message, http.StatusNotFound, underlying, requestId)
}

func NewForbiddenError(message string, underlying error, requestId string) *Error {
	return NewError(Forbidden, message, http.StatusForbidden, underlying, requestId)
}
