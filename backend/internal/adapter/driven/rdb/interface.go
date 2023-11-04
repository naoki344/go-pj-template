package rdbadapter

import (
	"errors"
)

var RdbErrCustomerNotFound = errors.New("CustomerNotFound")
var RdbErrUnexpected = errors.New("DBUnexpectedErrorOccurred")

type RdbInterface interface {
	GetCustomerByID(customerID int64) (*Customer, error)
}
