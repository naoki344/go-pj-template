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
	return &ogen.PostCreateCustomerOKHeaders{
		Response: ogen.PostCreateCustomerOK{
			ID:   int64(1),
			Name: string("三好 直紀"),
			NameKana: ogen.OptString{
				Value: "ミヨシ ナオキ",
				Set:   true,
			},
			Telephone:          "09011112222",
			Email:              "example@gmail.com",
			PersonInChargeName: "テスト名",
			PersonInChargeNameKana: ogen.OptString{
				Value: "テストメイ",
				Set:   true,
			},
			Address: ogen.PostCreateCustomerOKAddress{
				PostalCode: "8800301",
				PrefID:     1,
				Address1:   "宮崎市佐土原町",
				Address2:   "1-10-10",
			},
		},
	}, nil
}

func (n *EnAPIAdapter) PostSearchCustomer(ctx context.Context, req *ogen.PostSearchCustomerReq) (ogen.PostSearchCustomerRes, error) {
	return &ogen.PostSearchCustomerOKHeaders{
		Response: ogen.PostSearchCustomerOK{
			Page: ogen.PostSearchCustomerOKPage{
				Size:    10,
				Total:   10021,
				Current: 2,
			},
			Customers: []ogen.PostSearchCustomerOKCustomersItem{
				{
					ID:   int64(1),
					Name: string("三好 直紀"),
					NameKana: ogen.OptString{
						Value: "ミヨシ ナオキ",
						Set:   true,
					},
					Telephone:          "09011112222",
					Email:              "example@gmail.com",
					PersonInChargeName: "テスト名",
					PersonInChargeNameKana: ogen.OptString{
						Value: "テストメイ",
						Set:   true,
					},
					Address: ogen.PostSearchCustomerOKCustomersItemAddress{
						PostalCode: "8800301",
						PrefID:     1,
						Address1:   "宮崎市佐土原町",
						Address2:   "1-10-10",
					},
				},
				{
					ID:   int64(2),
					Name: string("大輝証券"),
					NameKana: ogen.OptString{
						Value: "ダイキショウケン",
						Set:   true,
					},
					Telephone:          "09011112222",
					Email:              "example@gmail.com",
					PersonInChargeName: "テスト名",
					PersonInChargeNameKana: ogen.OptString{
						Value: "テストメイ",
						Set:   true,
					},
					Address: ogen.PostSearchCustomerOKCustomersItemAddress{
						PostalCode: "8800301",
						PrefID:     1,
						Address1:   "宮崎市佐土原町",
						Address2:   "1-10-10",
					},
				},
			},
		},
	}, nil
}

func (n *EnAPIAdapter) PutModifyCustomerByID(ctx context.Context, req *ogen.PutModifyCustomerByIDReq, params ogen.PutModifyCustomerByIDParams) (ogen.PutModifyCustomerByIDRes, error) {

	err := n.customerAPI.UpdateByID()
	if err != nil {
		return CreateErrorPutByIDResponse(err), nil
	}
	return &ogen.PutModifyCustomerByIDOKHeaders{
		Response: ogen.PutModifyCustomerByIDOK{
			ID:   int64(1),
			Name: string("三好 直紀"),
			NameKana: ogen.OptString{
				Value: "ミヨシ ナオキ",
				Set:   true,
			},
			Telephone:          "09011112222",
			Email:              "example@gmail.com",
			PersonInChargeName: "テスト名",
			PersonInChargeNameKana: ogen.OptString{
				Value: "テストメイ",
				Set:   true,
			},
			Address: ogen.PutModifyCustomerByIDOKAddress{
				PostalCode: "8800301",
				PrefID:     1,
				Address1:   "宮崎市佐土原町",
				Address2:   "1-10-10",
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
			NameKana:               CreateOptString(res.NameKana),
			Telephone:              res.Telephone,
			Email:                  res.Email,
			PersonInChargeName:     res.PersonInChargeName,
			PersonInChargeNameKana: CreateOptString(res.PersonInChargeNameKana),
			Address: ogen.GetCustomerByIDOKAddress{
				PostalCode: res.Address1,
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

func CreateOptString(value *string) ogen.OptString {
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

func NewEnAPIAdapter(customerAPI *apiport.CustomerAPIPort) *EnAPIAdapter {
	return &EnAPIAdapter{
		customerAPI: customerAPI,
	}
}
