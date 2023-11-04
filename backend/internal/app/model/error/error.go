package errormodel


import (
	"errors"
)

var CustomerNotFound = errors.New("CustomerNotFound")
var UnexpectedError = errors.New("UnexpectedError")
