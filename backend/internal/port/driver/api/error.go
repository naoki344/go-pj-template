package apiport

import (
	"errors"
	"fmt"
)

type APICustomerNotFoundError struct {
	CustomerID CustomerID
}

func (err *APICustomerNotFoundError) Error() string {
	return fmt.Sprintf("CustomerNotFound [CustomerID=%d]", err.CustomerID)
}

var ErrUnexpected = errors.New("UnexpectedError")
