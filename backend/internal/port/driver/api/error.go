package apiport


import (
	"errors"
	"fmt"
)

type APIErrCustomerNotFound struct {
	customerid CustomerID
}

func (err *APIErrCustomerNotFound)Error() string {
	return	fmt.Sprintf("CustomerNotFound [CustomerID=%d]", err.customerid)
}

var UnexpectedError = errors.New("UnexpectedError")
