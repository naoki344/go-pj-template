package customerusecase

import (
	customermodel "github.com/g-stayfresh/en/backend/internal/domain/model/customer"
)


type GetCustomerByIDInterface interface {
	Run(customerID customermodel.ID) (*customermodel.Customer, error)
}
