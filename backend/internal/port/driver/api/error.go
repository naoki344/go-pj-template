package apiport

import (
	"fmt"

	"github.com/cockroachdb/errors"
)

type APICustomerNotFoundError struct {
	customerID CustomerID
	err        error
}

func (e *APICustomerNotFoundError) Error() string {
	return fmt.Sprintf("CustomerNotFound [CustomerID=%d]", e.customerID)
}

func (e *APICustomerNotFoundError) Unwrap() error {
	return e.err
}

func NewAPICustomerNotFoundError(err error, customerID CustomerID) error {
	return &APICustomerNotFoundError{
		err:        errors.WithStack(err),
		customerID: customerID,
	}
}

type APIUnexpectedError struct {
	err error
}

func (e *APIUnexpectedError) Error() string {
	return "UnexpectedError"
}

func (e *APIUnexpectedError) Unwrap() error {
	return e.err
}

func NewAPIUnexpectedError(err error) error {
	return &APIUnexpectedError{
		err: errors.WithStack(err),
	}
}
