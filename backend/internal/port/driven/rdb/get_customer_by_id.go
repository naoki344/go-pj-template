package rdbport


import (
	customermodel "github.com/g-stayfresh/en/backend/internal/app/model/customer"
	errormodel "github.com/g-stayfresh/en/backend/internal/app/model/error"
	rdbadapter "github.com/g-stayfresh/en/backend/internal/adapter/driven/rdb"
)


type GetCustomerByIDPort struct {
	rdb rdbadapter.RdbInterface
}

func NewGetCustomerByIDPort(rdb rdbadapter.RdbInterface) *GetCustomerByIDPort {
	return &GetCustomerByIDPort{rdb}
}

func (port *GetCustomerByIDPort) Get(customerId customermodel.CustomerID) (*customermodel.Customer, error){
	res, err := port.rdb.GetCustomerByID(int64(customerId))
	if err != nil {
		if err == rdbadapter.RdbErrCustomerNotFound {
			return nil, errormodel.CustomerNotFound
		}
		return nil, errormodel.UnexpectedError
	}
	customer := customermodel.Customer{
		ID: customermodel.CustomerID(res.ID),
		Title: res.Title,
		Content: res.Content,
	}
	return &customer, nil
}
