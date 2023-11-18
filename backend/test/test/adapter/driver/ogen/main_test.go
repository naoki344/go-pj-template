package ogenadapter

import (
	"context"
	"testing"

	"github.com/g-stayfresh/en/backend/api/lib/ogen"
	apiport "github.com/g-stayfresh/en/backend/internal/port/driver/api"
	"github.com/stretchr/testify/assert"
)

func TestEnAPIAdapter_PostCreateCustomer(t *testing.T) {
	type fields struct {
		customerAPI *apiport.CustomerAPIPort
	}
	type args struct {
		ctx context.Context
		req *ogen.PostCreateCustomerRequest
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      ogen.PostCreateCustomerRes
		assertion assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &EnAPIAdapter{
				customerAPI: tt.fields.customerAPI,
			}
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
		req *ogen.PostSearchCustomerRequest
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      ogen.PostSearchCustomerRes
		assertion assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &EnAPIAdapter{
				customerAPI: tt.fields.customerAPI,
			}
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
		req    *ogen.PutModifyCustomerByIDRequest
		params ogen.PutModifyCustomerByIDParams
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      ogen.PutModifyCustomerByIDRes
		assertion assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &EnAPIAdapter{
				customerAPI: tt.fields.customerAPI,
			}
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
		params ogen.GetCustomerByIDParams
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      ogen.GetCustomerByIDRes
		assertion assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &EnAPIAdapter{
				customerAPI: tt.fields.customerAPI,
			}
			got, err := n.GetCustomerByID(tt.args.ctx, tt.args.params)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_toCustomer(t *testing.T) {
	type args struct {
		customer *apiport.Customer
	}
	tests := []struct {
		name string
		args args
		want ogen.Customer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, toCustomer(tt.args.customer))
		})
	}
}

func Test_createCustomerSearchResponse(t *testing.T) {
	type args struct {
		page      apiport.PageResult
		customers []*apiport.Customer
	}
	tests := []struct {
		name string
		args args
		want *ogen.PostSearchCustomer200ResponseHeaders
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, createCustomerSearchResponse(tt.args.page, tt.args.customers))
		})
	}
}

func Test_createCustomerResponse(t *testing.T) {
	type args struct {
		customer *apiport.Customer
	}
	tests := []struct {
		name string
		args args
		want *ogen.CustomerHeaders
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, createCustomerResponse(tt.args.customer))
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
		want ogen.GetCustomerByIDRes
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, CreateErrorGetByIDResponse(tt.args.err))
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
		want ogen.PutModifyCustomerByIDRes
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, CreateErrorPutByIDResponse(tt.args.err))
		})
	}
}

func TestCreateErrorPutByIDResponseUnmatchID(t *testing.T) {
	tests := []struct {
		name string
		want ogen.PutModifyCustomerByIDRes
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, CreateErrorPutByIDResponseUnmatchID())
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
		want ogen.PostCreateCustomerRes
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, CreateErrorPostCreateCustomerResponse(tt.args.err))
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
		want ogen.PostSearchCustomerRes
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, CreateErrorPostSearchCustomerResponse(tt.args.err))
		})
	}
}

func Test_toOptString(t *testing.T) {
	type args struct {
		value *string
	}
	tests := []struct {
		name string
		args args
		want ogen.OptString
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, toOptString(tt.args.value))
		})
	}
}

func Test_getStringFromOptString(t *testing.T) {
	type args struct {
		optString ogen.OptString
	}
	tests := []struct {
		name string
		args args
		want *string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, getStringFromOptString(tt.args.optString))
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
		want *EnAPIAdapter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewEnAPIAdapter(tt.args.customerAPI))
		})
	}
}
