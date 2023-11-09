package apiport

import (
	"errors"
	"fmt"
)

type APICustomerNotFoundError struct {
	customerid CustomerID
}

func (err *APICustomerNotFoundError) Error() string {
	return fmt.Sprintf("CustomerNotFound [CustomerID=%d]", err.customerid)
}

var ErrUnexpected = errors.New("UnexpectedError")
