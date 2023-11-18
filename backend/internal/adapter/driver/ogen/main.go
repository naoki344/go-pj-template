package ogenadapter

import (
	"context"
	"errors"
	"log/slog"

	ogen "github.com/g-stayfresh/en/backend/internal/adapter/driver/ogenlib"
	apiport "github.com/g-stayfresh/en/backend/internal/port/driver/api"
)

type ErrorType string

const (
	ResourceNotFound    = ErrorType("ResourceNotFound")
	InternalServerError = ErrorType("InternalServerError")
)

type EnAPIAdapter struct {
	customerAPI *apiport.CustomerAPIPort
}

func (n *EnAPIAdapter) PostCreateCustomer(ctx context.Context, req *ogen.PostCreateCustomerRequest) (ogen.PostCreateCustomerRes, error) {
	portModel := &apiport.Customer{
		Name:                   req.Name,
		NameKana:               getStringFromOptString(req.NameKana),
		Telephone:              req.Telephone,
		Email:                  req.Email,
		PersonInChargeName:     req.PersonInChargeName,
		PersonInChargeNameKana: getStringFromOptString(req.PersonInChargeNameKana),
		PostalCode:             req.Address.PostalCode,
		PrefID:                 req.Address.PrefID,
		Address1:               req.Address.Address1,
		Address2:               req.Address.Address2,
	}
	res, err := n.customerAPI.CreateCustomer(portModel)
	if err != nil {
		return CreateErrorPostCreateCustomerResponse(err), nil
	}
	return createCustomerResponse(res), nil
}

func (n *EnAPIAdapter) PostSearchCustomer(ctx context.Context, req *ogen.PostSearchCustomerRequest) (ogen.PostSearchCustomerRes, error) {
	pageNumber := req.Pagination.Number
	pageSize := req.Pagination.Size
	res, err := n.customerAPI.SearchCustomer(
		pageNumber, pageSize, &apiport.SearchConditions{})
	if err != nil {
		return CreateErrorPostSearchCustomerResponse(err), nil
	}
	return createCustomerSearchResponse(res.Page, res.CustomerList), nil
}

func (n *EnAPIAdapter) PutModifyCustomerByID(ctx context.Context, req *ogen.PutModifyCustomerByIDRequest, params ogen.PutModifyCustomerByIDParams) (ogen.PutModifyCustomerByIDRes, error) {
	if params.CustomerID != req.ID {
		return CreateErrorPutByIDResponseUnmatchID(), nil
	}
	portModel := &apiport.Customer{
		ID:                     req.ID,
		Name:                   req.Name,
		NameKana:               getStringFromOptString(req.NameKana),
		Telephone:              req.Telephone,
		Email:                  req.Email,
		PersonInChargeName:     req.PersonInChargeName,
		PersonInChargeNameKana: getStringFromOptString(req.PersonInChargeNameKana),
		PostalCode:             req.Address.PostalCode,
		PrefID:                 req.Address.PrefID,
		Address1:               req.Address.Address1,
		Address2:               req.Address.Address2,
	}

	res, err := n.customerAPI.UpdateByID(portModel)
	if err != nil {
		return CreateErrorPutByIDResponse(err), nil
	}
	return createCustomerResponse(res), nil
}

func (n *EnAPIAdapter) GetCustomerByID(ctx context.Context, params ogen.GetCustomerByIDParams) (ogen.GetCustomerByIDRes, error) {
	res, err := n.customerAPI.GetByID(apiport.CustomerID(params.CustomerID))
	if err != nil {
		return CreateErrorGetByIDResponse(err), nil
	}
	return createCustomerResponse(res), nil
}

func toCustomer(customer *apiport.Customer) ogen.Customer {
	return ogen.Customer{
		ID:                     customer.ID,
		Name:                   customer.Name,
		NameKana:               toOptString(customer.NameKana),
		Telephone:              customer.Telephone,
		Email:                  customer.Email,
		PersonInChargeName:     customer.PersonInChargeName,
		PersonInChargeNameKana: toOptString(customer.PersonInChargeNameKana),
		Address: ogen.Address{
			PostalCode: customer.PostalCode,
			PrefID:     customer.PrefID,
			Address1:   customer.Address1,
			Address2:   customer.Address2,
		},
	}
}

func createCustomerSearchResponse(page apiport.PageResult, customers []*apiport.Customer) *ogen.PostSearchCustomer200ResponseHeaders {
	customersRes := make([]ogen.Customer, 0)
	for _, v := range customers {
		customersRes = append(customersRes, toCustomer(v))
	}
	as := string('*')
	return &ogen.PostSearchCustomer200ResponseHeaders{
		AccessControlAllowHeaders: toOptString(&as),
		AccessControlAllowMethods: toOptString(&as),
		AccessControlAllowOrigin:  toOptString(&as),
		Response: ogen.PostSearchCustomer200Response{
			Page: ogen.Page{
				Size:    page.Size,
				Total:   page.Total,
				Current: page.Current,
			},
			Customers: customersRes,
		},
	}
}

func createCustomerResponse(customer *apiport.Customer) *ogen.CustomerHeaders {
	as := string('*')
	return &ogen.CustomerHeaders{
		AccessControlAllowHeaders: toOptString(&as),
		AccessControlAllowMethods: toOptString(&as),
		AccessControlAllowOrigin:  toOptString(&as),
		Response:                  toCustomer(customer),
	}
}

func CreateErrorGetByIDResponse(err error) ogen.GetCustomerByIDRes {
	slog.Error("get customer error.", "err", err)
	as := string('*')
	var customerErr *apiport.APICustomerNotFoundError
	if errors.As(err, &customerErr) {
		return &ogen.GetCustomerByIDNotFound{
			AccessControlAllowHeaders: toOptString(&as),
			AccessControlAllowMethods: toOptString(&as),
			AccessControlAllowOrigin:  toOptString(&as),
			Response: ogen.ErrorModel{
				Type:    string(ResourceNotFound),
				Message: "customer not found.",
			},
		}
	}
	return &ogen.GetCustomerByIDInternalServerError{
		AccessControlAllowHeaders: toOptString(&as),
		AccessControlAllowMethods: toOptString(&as),
		AccessControlAllowOrigin:  toOptString(&as),
		Response: ogen.ErrorModel{
			Type:    string(InternalServerError),
			Message: "unexpected error has occurred.",
		},
	}
}

func CreateErrorPutByIDResponse(err error) ogen.PutModifyCustomerByIDRes {
	slog.Error("get customer error.", "err", err)
	as := string('*')
	var customerErr *apiport.APICustomerNotFoundError
	if errors.As(err, &customerErr) {
		return &ogen.PutModifyCustomerByIDNotFound{
			AccessControlAllowHeaders: toOptString(&as),
			AccessControlAllowMethods: toOptString(&as),
			AccessControlAllowOrigin:  toOptString(&as),
			Response: ogen.ErrorModel{
				Type:    string(ResourceNotFound),
				Message: "customer not found.",
			},
		}
	}
	return &ogen.PutModifyCustomerByIDInternalServerError{
		AccessControlAllowHeaders: toOptString(&as),
		AccessControlAllowMethods: toOptString(&as),
		AccessControlAllowOrigin:  toOptString(&as),
		Response: ogen.ErrorModel{
			Type:    string(InternalServerError),
			Message: "unexpected error has occurred.",
		},
	}
}

func CreateErrorPutByIDResponseUnmatchID() ogen.PutModifyCustomerByIDRes {
	as := string('*')
	return &ogen.PutModifyCustomerByIDNotFound{
		AccessControlAllowHeaders: toOptString(&as),
		AccessControlAllowMethods: toOptString(&as),
		AccessControlAllowOrigin:  toOptString(&as),
		Response: ogen.ErrorModel{
			Type:    string(ResourceNotFound),
			Message: "customer id unmatch.",
		},
	}
}

func CreateErrorPostCreateCustomerResponse(err error) ogen.PostCreateCustomerRes {
	slog.Error("get customer error.", "err", err)
	as := string('*')
	return &ogen.PostCreateCustomerInternalServerError{
		AccessControlAllowHeaders: toOptString(&as),
		AccessControlAllowMethods: toOptString(&as),
		AccessControlAllowOrigin:  toOptString(&as),
		Response: ogen.ErrorModel{
			Type:    string(InternalServerError),
			Message: "unexpected error has occurred.",
		},
	}
}

func CreateErrorPostSearchCustomerResponse(err error) ogen.PostSearchCustomerRes {
	slog.Error("get customer error.", "err", err)
	as := string('*')
	return &ogen.PostSearchCustomerInternalServerError{
		AccessControlAllowHeaders: toOptString(&as),
		AccessControlAllowMethods: toOptString(&as),
		AccessControlAllowOrigin:  toOptString(&as),
		Response: ogen.ErrorModel{
			Type:    string(InternalServerError),
			Message: "unexpected error has occurred.",
		},
	}
}

func toOptString(value *string) ogen.OptString {
	if value != nil {
		return ogen.OptString{
			Value: *value,
			Set:   true,
		}
	}
	var v string
	return ogen.OptString{
		Value: v,
		Set:   false,
	}
}

func getStringFromOptString(optString ogen.OptString) *string {
	if optString.Set {
		return &optString.Value
	}
	return nil
}

func NewEnAPIAdapter(customerAPI *apiport.CustomerAPIPort) *EnAPIAdapter {
	return &EnAPIAdapter{
		customerAPI: customerAPI,
	}
}
