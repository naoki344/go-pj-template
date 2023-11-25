package rdbport

import (
	"github.com/cockroachdb/errors"
	rdbadapter "github.com/g-stayfresh/en/backend/internal/adapter/driven/rdb"
	errormodel "github.com/g-stayfresh/en/backend/internal/domain/error"
	customermodel "github.com/g-stayfresh/en/backend/internal/domain/model/customer"
	pagemodel "github.com/g-stayfresh/en/backend/internal/domain/model/page"
)

type RdbPort struct {
	rdb rdbadapter.RdbInterface
}

func NewRdbPort(rdb rdbadapter.RdbInterface) *RdbPort {
	return &RdbPort{rdb}
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

func toModelCustomerList(customers []rdbadapter.Customer) *[]*customermodel.Customer {
	models := make([]*customermodel.Customer, 0)
	for i := range customers {
		models = append(models, toModelCustomer(&customers[i]))
	}
	return &models
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

func (port *RdbPort) CustomerCreate(customer *customermodel.Customer) (*customermodel.Customer, error) {
	ogenCustomer := toAdapterCustomer(customer)
	res, err := port.rdb.InsertCustomer(ogenCustomer)
	if err != nil {
		var notFoundErr *rdbadapter.RdbCustomerNotFoundError
		if errors.As(err, &notFoundErr) {
			return nil, errormodel.NewCustomerNotFoundError(err)
		}
		return nil, errormodel.NewUnexpectedError(err)
	}
	return toModelCustomer(res), nil
}

func (port *RdbPort) CustomerGet(customerID customermodel.ID) (*customermodel.Customer, error) {
	res, err := port.rdb.GetCustomerByID(int64(customerID))
	if err != nil {
		var notFoundErr *rdbadapter.RdbCustomerNotFoundError
		if errors.As(err, &notFoundErr) {
			return nil, errormodel.NewCustomerNotFoundError(err)
		}
		return nil, errormodel.NewUnexpectedError(err)
	}
	return toModelCustomer(res), nil
}

func (port *RdbPort) CustomerUpdate(customer *customermodel.Customer) error {
	ogenCustomer := toAdapterCustomer(customer)
	err := port.rdb.UpdateCustomerByID(ogenCustomer)
	if err != nil {
		var notFoundErr *rdbadapter.RdbCustomerNotFoundError
		if errors.As(err, &notFoundErr) {
			return errormodel.NewCustomerNotFoundError(err)
		}
		return errormodel.NewUnexpectedError(err)
	}
	return nil
}

func (port *RdbPort) CustomerSearch(pageNumber int64, pageSize int64, conditions *customermodel.SearchConditions) (*[]*customermodel.Customer, *pagemodel.PageResult, error) {
	res, err := port.rdb.SearchCustomer(pageNumber, pageSize, &rdbadapter.SearchConditions{})
	if err != nil {
		return nil, nil, errormodel.NewUnexpectedError(err)
	}
	return toModelCustomerList(res.CustomerList), &pagemodel.PageResult{
		Size:    pagemodel.Size(res.PageInfo.Size),
		Total:   pagemodel.Total(res.PageInfo.Total),
		Current: pagemodel.Current(res.PageInfo.Current),
	}, nil
}
