package apiport

import (
	"errors"

	errormodel "github.com/g-stayfresh/en/backend/internal/domain/error"
	customermodel "github.com/g-stayfresh/en/backend/internal/domain/model/customer"
	customerusecase "github.com/g-stayfresh/en/backend/internal/usecase/customer"
)

type CustomerID int64

type Customer struct {
	ID                     int64
	Name                   string
	NameKana               *string
	Telephone              string
	Email                  string
	PersonInChargeName     string
	PersonInChargeNameKana *string
	PostalCode             string
	PrefID                 int64
	Address1               string
	Address2               string
}

func toPortCustomer(customer *customermodel.Customer) *Customer {
	return &Customer{
		ID:                     int64(customer.ID),
		Name:                   string(customer.Name),
		NameKana:               customer.NameKana,
		Telephone:              string(customer.Telephone),
		Email:                  string(customer.Email),
		PersonInChargeName:     string(customer.PersonInChargeName),
		PersonInChargeNameKana: customer.PersonInChargeNameKana,
		PostalCode:             string(customer.Address.PostalCode),
		PrefID:                 int64(customer.Address.PrefID),
		Address1:               string(customer.Address.Address1),
		Address2:               string(customer.Address.Address2),
	}
}

func toModelCustomer(customer *Customer) *customermodel.Customer {
	return &customermodel.Customer{}
}

type CustomerAPIPort struct {
	usecase customerusecase.CustomerUsecaseInterface
}

func NewCustomerAPIPort(usecase customerusecase.CustomerUsecaseInterface) *CustomerAPIPort {
	return &CustomerAPIPort{usecase}
}

func (port *CustomerAPIPort) GetByID(customerID CustomerID) (*Customer, error) {
	res, err := port.usecase.GetByID(customermodel.ID(customerID))
	if err != nil {
		if errors.Is(err, errormodel.ErrCustomerNotFound) {
			return nil, &APICustomerNotFoundError{customerID}
		}
		return nil, ErrUnexpected
	}
	return toPortCustomer(res), nil
}

func (port *CustomerAPIPort) UpdateByID(customer *Customer) error {
	err := port.usecase.UpdateByID(toModelCustomer(customer))
	if err != nil {
		if errors.Is(err, errormodel.ErrCustomerNotFound) {
			return &APICustomerNotFoundError{CustomerID(customer.ID)}
		}
		return ErrUnexpected
	}
	return nil
}
