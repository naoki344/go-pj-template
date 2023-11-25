package rdbadapter

import (
	"fmt"

	"github.com/cockroachdb/errors"
)

const (
	RdbCustomerNotFoundMessage = "CustomerNotFound"
	RdbUnexpectedMessage       = "DBUnexpectedErrorOccurred"
)

type RdbCustomerNotFoundError struct {
	err error
}

func (e *RdbCustomerNotFoundError) Error() string {
	return RdbCustomerNotFoundMessage
}

func (e *RdbCustomerNotFoundError) Unwrap() error {
	return e.err
}

func NewRdbCustomerNotFoundError(err error, customerID int64) error {
	msg := fmt.Sprintf("customer_id = %d", customerID)
	err = errors.WithHint(err, msg)
	err = errors.WithMessage(err, msg)
	newErr := errors.WithStack(&RdbCustomerNotFoundError{
		err: err,
	})
	return newErr
}

type RdbUnexpectedError struct {
	err error
}

func (e *RdbUnexpectedError) Error() string {
	return RdbUnexpectedMessage
}

func (e *RdbUnexpectedError) Unwrap() error {
	return e.err
}

func NewRdbUnexpectedError(err error) error {
	withStack := errors.WithStack(err)
	return &RdbUnexpectedError{withStack}
}

type CustomerList []Customer

type SearchConditions struct{}

type PageInfo struct {
	Size    int64
	Total   int64
	Current int64
}

type CustomerSearchResult struct {
	CustomerList CustomerList
	PageInfo     PageInfo
}

type RdbInterface interface {
	GetCustomerByID(customerID int64) (*Customer, error)
	UpdateCustomerByID(customer *Customer) error
	InsertCustomer(customer *Customer) (*Customer, error)
	SearchCustomer(pageNumber int64, pageSize int64, conditions *SearchConditions) (*CustomerSearchResult, error)
}
