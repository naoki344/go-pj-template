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

type GetCustomerByIDAPIPort struct {
	usecase customerusecase.GetCustomerByIDInterface
}

func NewGetCustomerByIDAPIPort(usecase customerusecase.GetCustomerByIDInterface) *GetCustomerByIDAPIPort {
	return &GetCustomerByIDAPIPort{usecase}
}

func (port *GetCustomerByIDAPIPort) Run(customerID CustomerID) (*Customer, error) {
	res, err := port.usecase.Run(customermodel.ID(customerID))
	if err != nil {
		if errors.Is(err, errormodel.ErrCustomerNotFound) {
			return nil, &APICustomerNotFoundError{customerID}
		}
		return nil, ErrUnexpected
	}
	return &Customer{
		ID:                     int64(res.ID),
		Name:                   string(res.Name),
		NameKana:               res.NameKana,
		Telephone:              string(res.Telephone),
		Email:                  string(res.Email),
		PersonInChargeName:     string(res.PersonInChargeName),
		PersonInChargeNameKana: res.PersonInChargeNameKana,
		PostalCode:             string(res.Address.PostalCode),
		PrefID:                 int64(res.Address.PrefID),
		Address1:               string(res.Address.Address1),
		Address2:               string(res.Address.Address2),
	}, nil
}
