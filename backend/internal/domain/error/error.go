package errormodel

import (
	"github.com/cockroachdb/errors"
)

const (
	CustomerNotFoundMessage = "CustomerNotFound"
	UnexpectedMessage       = "UnexpectedError"
)

type CustomerNotFoundError struct {
	err error
}

func (e *CustomerNotFoundError) Error() string {
	return CustomerNotFoundMessage
}

func (e *CustomerNotFoundError) Unwrap() error {
	return e.err
}

func NewCustomerNotFoundError(err error) error {
	return &CustomerNotFoundError{errors.WithStack(err)}
}

type UnexpectedError struct {
	err error
}

func (e *UnexpectedError) Error() string {
	return UnexpectedMessage
}

func (e *UnexpectedError) Unwrap() error {
	return e.err
}

func NewUnexpectedError(err error) error {
	return &UnexpectedError{errors.WithStack(err)}
}
