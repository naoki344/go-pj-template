package customerusecase

import (
	customermodel "github.com/g-stayfresh/en/backend/internal/app/model/customer"
	rdbport "github.com/g-stayfresh/en/backend/internal/port/driven/rdb"
)

type GetCustomerByIDUsecase struct {
	port rdbport.GetCustomerByIDPortInterface
}

func NewGetCustomerByIDUsecase (port rdbport.GetCustomerByIDPortInterface) *GetCustomerByIDUsecase{
	return &GetCustomerByIDUsecase{port}
}


func (usecase *GetCustomerByIDUsecase) Run(customerId customermodel.CustomerID) (*customermodel.Customer, error) {
	return usecase.port.Get(customerId)
}
