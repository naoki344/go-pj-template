package rdbport

import (
	"errors"

	rdbadapter "github.com/g-stayfresh/en/backend/internal/adapter/driven/rdb"
	errormodel "github.com/g-stayfresh/en/backend/internal/domain/error"
	customermodel "github.com/g-stayfresh/en/backend/internal/domain/model/customer"
)

type GetCustomerByIDPort struct {
	rdb rdbadapter.RdbInterface
}

func NewGetCustomerByIDPort(rdb rdbadapter.RdbInterface) *GetCustomerByIDPort {
	return &GetCustomerByIDPort{rdb}
}

func (port *GetCustomerByIDPort) Get(customerID customermodel.ID) (*customermodel.Customer, error) {
	res, err := port.rdb.GetCustomerByID(int64(customerID))
	if err != nil {
		if errors.Is(err, rdbadapter.ErrRdbCustomerNotFound) {
			return nil, errormodel.ErrCustomerNotFound
		}
		return nil, errormodel.ErrUnexpectedError
	}
	return &customermodel.Customer{
		ID:        customermodel.ID(res.ID),
		Name:      customermodel.Name(res.Name),
		NameKana:  customermodel.NameKana(res.NameKana),
		Telephone: customermodel.Telephone(res.Telephone),
		Email:     customermodel.Email(res.Email),
		PersonInChargeName: customermodel.PersonInChargeName(
			res.PersonInChargeName),
		PersonInChargeNameKana: customermodel.PersonInChargeNameKana(res.PersonInChargeNameKana),
		Address: customermodel.Address{
			PostalCode: customermodel.PostalCode(res.PostalCode),
			PrefID:     customermodel.PrefID(res.PrefID),
			Address1:   customermodel.Address1(res.Address1),
			Address2:   customermodel.Address2(res.Address2),
		},
	}, nil
}
