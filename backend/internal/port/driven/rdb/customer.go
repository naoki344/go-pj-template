package rdbport

import (
	"errors"

	rdbadapter "github.com/g-stayfresh/en/backend/internal/adapter/driven/rdb"
	errormodel "github.com/g-stayfresh/en/backend/internal/domain/error"
	customermodel "github.com/g-stayfresh/en/backend/internal/domain/model/customer"
)

type CustomerRdbPort struct {
	rdb rdbadapter.RdbInterface
}

func NewCustomerRdbPort(rdb rdbadapter.RdbInterface) *CustomerRdbPort {
	return &CustomerRdbPort{rdb}
}

func toModelCustomer(customer *rdbadapter.Customer) *customermodel.Customer {
	return &customermodel.Customer{
		ID:        customermodel.ID(customer.ID),
		Name:      customermodel.Name(customer.Name),
		NameKana:  customermodel.NameKana(customer.NameKana),
		Telephone: customermodel.Telephone(customer.Telephone),
		Email:     customermodel.Email(customer.Email),
		PersonInChargeName: customermodel.PersonInChargeName(
			customer.PersonInChargeName),
		PersonInChargeNameKana: customermodel.PersonInChargeNameKana(customer.PersonInChargeNameKana),
		Address: customermodel.Address{
			PostalCode: customermodel.PostalCode(customer.PostalCode),
			PrefID:     customermodel.PrefID(customer.PrefID),
			Address1:   customermodel.Address1(customer.Address1),
			Address2:   customermodel.Address2(customer.Address2),
		},
	}
}

func toAdapterCustomer(customer *customermodel.Customer) *rdbadapter.Customer {
	return &rdbadapter.Customer{
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

func (port *CustomerRdbPort) Get(customerID customermodel.ID) (*customermodel.Customer, error) {
	res, err := port.rdb.GetCustomerByID(int64(customerID))
	if err != nil {
		if errors.Is(err, rdbadapter.ErrRdbCustomerNotFound) {
			return nil, errormodel.ErrCustomerNotFound
		}
		return nil, errormodel.ErrUnexpectedError
	}
	return toModelCustomer(res), nil
}

func (port *CustomerRdbPort) Update(customer *customermodel.Customer) error {
	ogenCustomer := toAdapterCustomer(customer)
	err := port.rdb.UpdateCustomerByID(ogenCustomer)
	if err != nil {
		if errors.Is(err, rdbadapter.ErrRdbCustomerNotFound) {
			return errormodel.ErrCustomerNotFound
		}
		return errormodel.ErrUnexpectedError
	}
	return nil
}
