package customerusecase

import (
	customermodel "github.com/g-stayfresh/en/backend/internal/domain/model/customer"
)

type CustomerUsecaseInterface interface {
	GetByID(customerID customermodel.ID) (*customermodel.Customer, error)
	UpdateByID(customer *customermodel.Customer) error
}
