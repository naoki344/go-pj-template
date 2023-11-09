package errormodel

import (
	"errors"
)

var (
	ErrCustomerNotFound = errors.New("CustomerNotFound")
	ErrUnexpectedError  = errors.New("UnexpectedError")
)
