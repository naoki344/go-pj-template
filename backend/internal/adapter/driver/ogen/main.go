package ogenadapter

import (
	"context"
	"errors"
	"log/slog"

	"github.com/g-stayfresh/en/backend/api/lib/ogen"
	apiport "github.com/g-stayfresh/en/backend/internal/port/driver/api"
)

type EnAPIAdapter struct {
	customerAPI *apiport.CustomerAPIPort
}

func (n *EnAPIAdapter) PostCreateCustomer(ctx context.Context, req *ogen.PostCreateCustomerReq) (ogen.PostCreateCustomerRes, error) {
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
	return &ogen.PostCreateCustomerOKHeaders{
		Response: ogen.PostCreateCustomerOK{
			ID:                     res.ID,
			Name:                   res.Name,
			NameKana:               toOptString(res.NameKana),
			Telephone:              res.Telephone,
			Email:                  res.Email,
			PersonInChargeName:     res.PersonInChargeName,
			PersonInChargeNameKana: toOptString(res.PersonInChargeNameKana),
			Address: ogen.PostCreateCustomerOKAddress{
				PostalCode: res.PostalCode,
				PrefID:     res.PrefID,
				Address1:   res.Address1,
				Address2:   res.Address2,
			},
		},
	}, nil
}

func (n *EnAPIAdapter) PostSearchCustomer(ctx context.Context, req *ogen.PostSearchCustomerReq) (ogen.PostSearchCustomerRes, error) {
	pageNumber := req.Pagination.Number
	pageSize := req.Pagination.Size
	res, err := n.customerAPI.SearchCustomer(
		pageNumber, pageSize, &apiport.SearchConditions{})
	if err != nil {
		return CreateErrorPostSearchCustomerResponse(err), nil
	}
	customers := make([]ogen.PostSearchCustomerOKCustomersItem, 0)
	for _, v := range res.CustomerList {
		item := ogen.PostSearchCustomerOKCustomersItem{
			ID:                     v.ID,
			Name:                   v.Name,
			NameKana:               toOptString(v.NameKana),
			Telephone:              v.Telephone,
			Email:                  v.Email,
			PersonInChargeName:     v.PersonInChargeName,
			PersonInChargeNameKana: toOptString(v.PersonInChargeNameKana),
			Address: ogen.PostSearchCustomerOKCustomersItemAddress{
				PostalCode: v.PostalCode,
				PrefID:     v.PrefID,
				Address1:   v.Address1,
				Address2:   v.Address2,
			},
		}
		customers = append(customers, item)
	}
	return &ogen.PostSearchCustomerOKHeaders{
		Response: ogen.PostSearchCustomerOK{
			Page: ogen.PostSearchCustomerOKPage{
				Size:    res.Page.Size,
				Total:   res.Page.Total,
				Current: res.Page.Current,
			},
			Customers: customers,
		},
	}, nil
}

func (n *EnAPIAdapter) PutModifyCustomerByID(ctx context.Context, req *ogen.PutModifyCustomerByIDReq, params ogen.PutModifyCustomerByIDParams) (ogen.PutModifyCustomerByIDRes, error) {
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
	return &ogen.PutModifyCustomerByIDOKHeaders{
		Response: ogen.PutModifyCustomerByIDOK{
			ID:                     res.ID,
			Name:                   res.Name,
			NameKana:               toOptString(res.NameKana),
			Telephone:              res.Telephone,
			Email:                  res.Email,
			PersonInChargeName:     res.PersonInChargeName,
			PersonInChargeNameKana: toOptString(res.PersonInChargeNameKana),
			Address: ogen.PutModifyCustomerByIDOKAddress{
				PostalCode: res.PostalCode,
				PrefID:     res.PrefID,
				Address1:   res.Address1,
				Address2:   res.Address2,
			},
		},
	}, nil
}

func (n *EnAPIAdapter) GetCustomerByID(ctx context.Context, params ogen.GetCustomerByIDParams) (ogen.GetCustomerByIDRes, error) {
	res, err := n.customerAPI.GetByID(apiport.CustomerID(params.CustomerID))
	if err != nil {
		return CreateErrorGetByIDResponse(err), nil
	}
	return &ogen.GetCustomerByIDOKHeaders{
		Response: ogen.GetCustomerByIDOK{
			ID:                     res.ID,
			Name:                   res.Name,
			NameKana:               toOptString(res.NameKana),
			Telephone:              res.Telephone,
			Email:                  res.Email,
			PersonInChargeName:     res.PersonInChargeName,
			PersonInChargeNameKana: toOptString(res.PersonInChargeNameKana),
			Address: ogen.GetCustomerByIDOKAddress{
				PostalCode: res.PostalCode,
				PrefID:     res.PrefID,
				Address1:   res.Address1,
				Address2:   res.Address2,
			},
		},
	}, nil
}

func CreateErrorGetByIDResponse(err error) ogen.GetCustomerByIDRes {
	slog.Error("get customer error.", "err", err)
	var customerErr *apiport.APICustomerNotFoundError
	if errors.As(err, &customerErr) {
		return &ogen.GetCustomerByIDNotFoundHeaders{
			Response: ogen.GetCustomerByIDNotFound{
				Type:    "ResourceNotFound",
				Message: "aaaaaaaaaaaaaaa",
			},
		}
	}
	return &ogen.GetCustomerByIDInternalServerErrorHeaders{
		Response: ogen.GetCustomerByIDInternalServerError{
			Type:    "InternalServerError",
			Message: "aaaaaaaaaaaaaaa",
		},
	}
}

func CreateErrorPutByIDResponse(err error) ogen.PutModifyCustomerByIDRes {
	slog.Error("get customer error.", "err", err)
	var customerErr *apiport.APICustomerNotFoundError
	if errors.As(err, &customerErr) {
		return &ogen.PutModifyCustomerByIDBadRequestHeaders{
			Response: ogen.PutModifyCustomerByIDBadRequest{
				Type:    "ResourceNotFound",
				Message: "aaaaaaaaaaaaaaa",
			},
		}
	}
	return &ogen.PutModifyCustomerByIDInternalServerErrorHeaders{
		Response: ogen.PutModifyCustomerByIDInternalServerError{
			Type:    "InternalServerError",
			Message: "aaaaaaaaaaaaaaa",
		},
	}
}

func CreateErrorPostCreateCustomerResponse(err error) ogen.PostCreateCustomerRes {
	slog.Error("get customer error.", "err", err)
	var customerErr *apiport.APICustomerNotFoundError
	if errors.As(err, &customerErr) {
		return &ogen.PostCreateCustomerNotFoundHeaders{
			Response: ogen.PostCreateCustomerNotFound{
				Type:    "ResourceNotFound",
				Message: "aaaaaaaaaaaaaaa",
			},
		}
	}
	return &ogen.PostCreateCustomerInternalServerErrorHeaders{
		Response: ogen.PostCreateCustomerInternalServerError{
			Type:    "InternalServerError",
			Message: "aaaaaaaaaaaaaaa",
		},
	}
}

func CreateErrorPostSearchCustomerResponse(err error) ogen.PostSearchCustomerRes {
	slog.Error("get customer error.", "err", err)
	return &ogen.PostSearchCustomerInternalServerErrorHeaders{
		Response: ogen.PostSearchCustomerInternalServerError{
			Type:    "InternalServerError",
			Message: "aaaaaaaaaaaaaaa",
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
