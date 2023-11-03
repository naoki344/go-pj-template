package rdbport


import (
	customermodel "github.com/g-stayfresh/en/backend/internal/app/model/customer"
)


type GetCustomerByIDPortInterface interface {
	Get(customerId customermodel.CustomerID) (customermodel.Customer, error)
}
