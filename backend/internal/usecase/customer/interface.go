package customerusecase

import (
	customermodel "github.com/g-stayfresh/en/backend/internal/app/model/customer"
)


type GetCustomerByIDInterface interface {
	Run(customerID customermodel.CustomerID) (*customermodel.Customer, error)
}
