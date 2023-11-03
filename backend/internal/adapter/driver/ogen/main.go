/*
Package yourpackage does something interesting.
*/
package ogenadapter

import (
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

func (n *EnAPIAdapter) GetCustomerByID(ctx context.Context, params ogen.GetCustomerByIDParams) (*ogen.GetCustomerByIDOK, error) {
	res, _ := n.getCustomerByID.Run(apiport.CustomerID(params.CustomerID))
	return &ogen.GetCustomerByIDOK{
		ID:      int64(res.ID),
		Title:   res.Title,
		Content: res.Content,
	}, nil
}


func NewEnAPIAdapter(getCustomerByID *apiport.GetCustomerByIDAPIPort) *EnAPIAdapter
