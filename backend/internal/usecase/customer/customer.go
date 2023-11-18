package customerusecase

import (
	customermodel "github.com/g-stayfresh/en/backend/internal/domain/model/customer"
	pagemodel "github.com/g-stayfresh/en/backend/internal/domain/model/page"
	rdbport "github.com/g-stayfresh/en/backend/internal/port/driven/rdb"
)

type CustomerUsecase struct {
	port rdbport.RdbPortInterface
}

func NewCustomerUsecase(port rdbport.RdbPortInterface) *CustomerUsecase {
	return &CustomerUsecase{port}
}

func (usecase *CustomerUsecase) GetByID(customerID customermodel.ID) (*customermodel.Customer, error) {
	return usecase.port.CustomerGet(customerID) //nolint:wrapcheck
}

func (usecase *CustomerUsecase) UpdateByID(customer *customermodel.Customer) (*customermodel.Customer, error) {
	err := usecase.port.CustomerUpdate(customer)
	if err != nil {
		return nil, err //nolint:wrapcheck
	}
	newCustomer, err := usecase.port.CustomerGet(customer.ID)
	if err != nil {
		return nil, err //nolint:wrapcheck
	}
	return newCustomer, nil
}

func (usecase *CustomerUsecase) Create(customer *customermodel.Customer) (*customermodel.Customer, error) {
	return usecase.port.CustomerCreate(customer) //nolint:wrapcheck
}

func (usecase *CustomerUsecase) Search(pageNumber int64, pageSize int64, conditions *customermodel.SearchConditions) (*[]*customermodel.Customer, *pagemodel.PageResult, error) {
	return usecase.port.CustomerSearch(pageNumber, pageSize, conditions) //nolint:wrapcheck
}
