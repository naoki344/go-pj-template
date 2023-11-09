package customerusecase

import (
	customermodel "github.com/g-stayfresh/en/backend/internal/domain/model/customer"
	rdbport "github.com/g-stayfresh/en/backend/internal/port/driven/rdb"
)

type GetCustomerByIDUsecase struct {
	port rdbport.GetCustomerByIDPortInterface
}

func NewGetCustomerByIDUsecase(port rdbport.GetCustomerByIDPortInterface) *GetCustomerByIDUsecase {
	return &GetCustomerByIDUsecase{port}
}

func (usecase *GetCustomerByIDUsecase) Run(customerID customermodel.ID) (*customermodel.Customer, error) {
	return usecase.port.Get(customerID) //nolint:wrapcheck
}
