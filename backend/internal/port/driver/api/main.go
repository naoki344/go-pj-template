package apiport

import (
	"errors"

	errormodel "github.com/g-stayfresh/en/backend/internal/domain/error"
	customermodel "github.com/g-stayfresh/en/backend/internal/domain/model/customer"
	pagemodel "github.com/g-stayfresh/en/backend/internal/domain/model/page"
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

type SearchConditions struct{}

type PageResult struct {
	Size    int64
	Total   int64
	Current int64
}

type CustomerSearchResult struct {
	CustomerList []*Customer
	Page         PageResult
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

func toCustomerSearchResult(customers *[]*customermodel.Customer, page *pagemodel.PageResult) *CustomerSearchResult {
	models := make([]*Customer, 0)
	for _, v := range *customers {
		models = append(models, toPortCustomer(v))
	}
	return &CustomerSearchResult{
		CustomerList: models,
		Page: PageResult{
			Size:    int64(page.Size),
			Total:   int64(page.Total),
			Current: int64(page.Current),
		},
	}
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

func (port *CustomerAPIPort) UpdateByID(customer *Customer) (*Customer, error) {
	res, err := port.usecase.UpdateByID(toModelCustomer(customer))
	if err != nil {
		if errors.Is(err, errormodel.ErrCustomerNotFound) {
			return nil, &APICustomerNotFoundError{CustomerID(customer.ID)}
		}
		return nil, ErrUnexpected
	}
	return toPortCustomer(res), nil
}

func (port *CustomerAPIPort) CreateCustomer(customer *Customer) (*Customer, error) {
	res, err := port.usecase.Create(toModelCustomer(customer))
	if err != nil {
		if errors.Is(err, errormodel.ErrCustomerNotFound) {
			return nil, &APICustomerNotFoundError{CustomerID(customer.ID)}
		}
		return nil, ErrUnexpected
	}
	return toPortCustomer(res), nil
}

func (port *CustomerAPIPort) SearchCustomer(pageNumber int64, pageSize int64, conditions *SearchConditions) (*CustomerSearchResult, error) {
	customers, page, err := port.usecase.Search(pageNumber, pageSize, &customermodel.SearchConditions{})
	if err != nil {
		return nil, ErrUnexpected
	}
	return toCustomerSearchResult(customers, page), nil
}
