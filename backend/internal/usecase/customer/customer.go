package customerusecase

import (
	customermodel "github.com/g-stayfresh/en/backend/internal/domain/model/customer"
	rdbport "github.com/g-stayfresh/en/backend/internal/port/driven/rdb"
)

type CustomerUsecase struct {
	port rdbport.CustomerRdbPortInterface
}

func NewCustomerUsecase(port rdbport.CustomerRdbPortInterface) *CustomerUsecase {
	return &CustomerUsecase{port}
}

func (usecase *CustomerUsecase) GetByID(customerID customermodel.ID) (*customermodel.Customer, error) {
	return usecase.port.Get(customerID) //nolint:wrapcheck
}

func (usecase *CustomerUsecase) UpdateByID(customer *customermodel.Customer) error {
	return usecase.port.Update(customer) //nolint:wrapcheck
}
