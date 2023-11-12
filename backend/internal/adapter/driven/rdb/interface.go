package rdbadapter

import (
	"errors"
)

var (
	ErrRdbCustomerNotFound = errors.New("CustomerNotFound")
	ErrRdbUnexpected       = errors.New("DBUnexpectedErrorOccurred")
)

type RdbInterface interface {
	GetCustomerByID(customerID int64) (*Customer, error)
	UpdateCustomerByID(customer *Customer) error
}
