package rdbport


import (
	customermodel "github.com/g-stayfresh/en/backend/internal/domain/model/customer"
)


type GetCustomerByIDPortInterface interface {
	Get(customerId customermodel.ID) (*customermodel.Customer, error)
}
