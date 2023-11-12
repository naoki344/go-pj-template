package rdbport

import (
	customermodel "github.com/g-stayfresh/en/backend/internal/domain/model/customer"
)

type CustomerRdbPortInterface interface {
	Get(customerID customermodel.ID) (*customermodel.Customer, error)
	Update(customer *customermodel.Customer) error
}
