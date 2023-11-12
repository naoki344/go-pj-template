package rdbadapter

import (
	"errors"
)

var (
	ErrRdbCustomerNotFound = errors.New("CustomerNotFound")
	ErrRdbUnexpected       = errors.New("DBUnexpectedErrorOccurred")
)

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
