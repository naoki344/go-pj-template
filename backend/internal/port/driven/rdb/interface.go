package rdbport

import (
	customermodel "github.com/g-stayfresh/en/backend/internal/domain/model/customer"
	pagemodel "github.com/g-stayfresh/en/backend/internal/domain/model/page"
)

type CustomerRdbPortInterface interface {
	Get(customerID customermodel.ID) (*customermodel.Customer, error)
	Update(customer *customermodel.Customer) error
	Create(customer *customermodel.Customer) (*customermodel.Customer, error)
	Search(pageNumber int64, pageSize int64, conditions *customermodel.SearchConditions) (*[]*customermodel.Customer, *pagemodel.PageResult, error)
}
