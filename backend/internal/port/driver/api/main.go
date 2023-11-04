package apiport


import (
	customerusecase "github.com/g-stayfresh/en/backend/internal/usecase/customer"
	customermodel "github.com/g-stayfresh/en/backend/internal/app/model/customer"
	errormodel "github.com/g-stayfresh/en/backend/internal/app/model/error"
)


type CustomerID int64

type Customer struct {
	ID      CustomerID
	Title   string
	Content string
}

type GetCustomerByIDAPIPort struct {
	usecase customerusecase.GetCustomerByIDInterface
}

func NewGetCustomerByIDAPIPort(usecase customerusecase.GetCustomerByIDInterface) *GetCustomerByIDAPIPort {
	return &GetCustomerByIDAPIPort{usecase}
}

func (port *GetCustomerByIDAPIPort) Run(customerId CustomerID) (*Customer, error){
	res, err := port.usecase.Run(customermodel.CustomerID(customerId))
	if err != nil {
		if err == errormodel.CustomerNotFound {
			return nil, &APIErrCustomerNotFound{customerId}
		}
		return nil, errormodel.UnexpectedError
	}
	return &Customer{
		ID: CustomerID(res.ID),
		Title: res.Title,
		Content: res.Content,
	}, nil
}
