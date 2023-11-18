package rdbport

import (
	customermodel "github.com/g-stayfresh/en/backend/internal/domain/model/customer"
	pagemodel "github.com/g-stayfresh/en/backend/internal/domain/model/page"
)

type RdbPortInterface interface {
	CustomerGet(customerID customermodel.ID) (*customermodel.Customer, error)
	CustomerUpdate(customer *customermodel.Customer) error
	CustomerCreate(customer *customermodel.Customer) (*customermodel.Customer, error)
	CustomerSearch(pageNumber int64, pageSize int64, conditions *customermodel.SearchConditions) (*[]*customermodel.Customer, *pagemodel.PageResult, error)
}
