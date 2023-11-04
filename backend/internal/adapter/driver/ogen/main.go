/*
Package yourpackage does something interesting.
*/
package ogenadapter

import (
	"errors"
	"github.com/g-stayfresh/en/backend/api/lib/ogen"
	apiport "github.com/g-stayfresh/en/backend/internal/port/driver/api"
	"log/slog"
	"context"
)


type EnAPIAdapter struct {
	getCustomerByID *apiport.GetCustomerByIDAPIPort
}

func (n *EnAPIAdapter) PostCreateCustomer(ctx context.Context, req *ogen.PostCreateCustomerReq) (*ogen.PostCreateCustomerOK, error) {
	return &ogen.PostCreateCustomerOK{}, nil
}

func (n *EnAPIAdapter) GetCustomerByID(ctx context.Context, params ogen.GetCustomerByIDParams) (ogen.GetCustomerByIDRes, error) {
	res, err := n.getCustomerByID.Run(apiport.CustomerID(params.CustomerID))
	if err != nil {
		return CreateErrorResponse(err), nil
	}
	return &ogen.GetCustomerByIDOK{
		ID:      int64(res.ID),
		Title:   res.Title,
		Content: res.Content,
	}, nil
}

func CreateErrorResponse(err error) ogen.GetCustomerByIDRes{
	var customerErr *apiport.APIErrCustomerNotFound
	if errors.As(err, &customerErr) {
		return &ogen.GetCustomerByIDNotFound{
			Type: "ResourceNotFound",
			Message: "aaaaaaaaaaaaaaa",

		}
	}
	return &ogen.GetCustomerByIDInternalServerError{
		Type: "SystemError",
		Message: "unexpected error occurred.",

	}
}


func NewEnAPIAdapter(getCustomerByID *apiport.GetCustomerByIDAPIPort) *EnAPIAdapter {
	slog.Error("create APIAdapter")
	return &EnAPIAdapter{
		getCustomerByID: getCustomerByID,
	}
}
