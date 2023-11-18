package ogenadapter_test

import (
	"context"
	"testing"

	ogen "github.com/g-stayfresh/en/backend/internal/adapter/driver/ogen"
	ogenlib "github.com/g-stayfresh/en/backend/internal/adapter/driver/ogenlib"
	apiport "github.com/g-stayfresh/en/backend/internal/port/driver/api"
	"github.com/stretchr/testify/assert"
)

func TestEnAPIAdapter_PostCreateCustomer(t *testing.T) {
	type fields struct {
		customerAPI *apiport.CustomerAPIPort
	}
	type args struct {
		ctx context.Context
		req *ogenlib.PostCreateCustomerRequest
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      ogenlib.PostCreateCustomerRes
		assertion assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := ogen.NewEnAPIAdapter(tt.fields.customerAPI)
			got, err := n.PostCreateCustomer(tt.args.ctx, tt.args.req)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestEnAPIAdapter_PostSearchCustomer(t *testing.T) {
	type fields struct {
		customerAPI *apiport.CustomerAPIPort
	}
	type args struct {
		ctx context.Context
		req *ogenlib.PostSearchCustomerRequest
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      ogenlib.PostSearchCustomerRes
		assertion assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := ogen.NewEnAPIAdapter(tt.fields.customerAPI)
			got, err := n.PostSearchCustomer(tt.args.ctx, tt.args.req)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestEnAPIAdapter_PutModifyCustomerByID(t *testing.T) {
	type fields struct {
		customerAPI *apiport.CustomerAPIPort
	}
	type args struct {
		ctx    context.Context
		req    *ogenlib.PutModifyCustomerByIDRequest
		params ogenlib.PutModifyCustomerByIDParams
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      ogenlib.PutModifyCustomerByIDRes
		assertion assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := ogen.NewEnAPIAdapter(tt.fields.customerAPI)
			got, err := n.PutModifyCustomerByID(tt.args.ctx, tt.args.req, tt.args.params)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestEnAPIAdapter_GetCustomerByID(t *testing.T) {
	type fields struct {
		customerAPI *apiport.CustomerAPIPort
	}
	type args struct {
		ctx    context.Context
		params ogenlib.GetCustomerByIDParams
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      ogenlib.GetCustomerByIDRes
		assertion assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := ogen.NewEnAPIAdapter(tt.fields.customerAPI)
			got, err := n.GetCustomerByID(tt.args.ctx, tt.args.params)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCreateErrorGetByIDResponse(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want ogenlib.GetCustomerByIDRes
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, ogen.CreateErrorGetByIDResponse(tt.args.err))
		})
	}
}

func TestCreateErrorPutByIDResponse(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want ogenlib.PutModifyCustomerByIDRes
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, ogen.CreateErrorPutByIDResponse(tt.args.err))
		})
	}
}

func TestCreateErrorPutByIDResponseUnmatchID(t *testing.T) {
	tests := []struct {
		name string
		want ogenlib.PutModifyCustomerByIDRes
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, ogen.CreateErrorPutByIDResponseUnmatchID())
		})
	}
}

func TestCreateErrorPostCreateCustomerResponse(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want ogenlib.PostCreateCustomerRes
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, ogen.CreateErrorPostCreateCustomerResponse(tt.args.err))
		})
	}
}

func TestCreateErrorPostSearchCustomerResponse(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want ogenlib.PostSearchCustomerRes
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, ogen.CreateErrorPostSearchCustomerResponse(tt.args.err))
		})
	}
}

func TestNewEnAPIAdapter(t *testing.T) {
	type args struct {
		customerAPI *apiport.CustomerAPIPort
	}
	tests := []struct {
		name string
		args args
		want *ogen.EnAPIAdapter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, ogen.NewEnAPIAdapter(tt.args.customerAPI))
		})
	}
}
