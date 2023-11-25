package ogenadapter_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/cockroachdb/errors"
	ogen "github.com/g-stayfresh/en/backend/internal/adapter/driver/ogen"
	ogenlib "github.com/g-stayfresh/en/backend/internal/adapter/driver/ogenlib"
	apiport "github.com/g-stayfresh/en/backend/internal/port/driver/api"
	apiportMock "github.com/g-stayfresh/en/backend/test/mock/port/driver/api"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestEnAPIAdapter_PostCreateCustomer(t *testing.T) {
	optString := "テスト名"
	req := ogenlib.PostCreateCustomerRequest{
		Name:                   "testname",
		NameKana:               toOptString(&optString),
		Telephone:              "09100001111",
		Email:                  "example.com",
		PersonInChargeName:     "person",
		PersonInChargeNameKana: toOptString(&optString),
		Address: ogenlib.Address{
			PostalCode: "8891111",
			PrefID:     1,
			Address1:   "宮崎市",
			Address2:   "佐土原",
		},
	}
	portModelReq := &apiport.Customer{
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
	portModelRes := &apiport.Customer{
		ID:                     12,
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
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	portMock := apiportMock.NewMockCustomerAPIPortInterface(ctrl)
	portMock.EXPECT().CreateCustomer(portModelReq).Return(portModelRes, nil)
	portMock2 := apiportMock.NewMockCustomerAPIPortInterface(ctrl)
	mockErr := apiport.NewAPIUnexpectedError(errors.New("test error"))
	portMock2.EXPECT().CreateCustomer(portModelReq).Return(nil, mockErr)
	ctx := context.Background()
	as := "*"
	expect := &ogenlib.CustomerHeaders{
		AccessControlAllowHeaders: toOptString(&as),
		AccessControlAllowMethods: toOptString(&as),
		AccessControlAllowOrigin:  toOptString(&as),
		Response: ogenlib.Customer{
			ID:                     12,
			Name:                   "testname",
			NameKana:               toOptString(&optString),
			Telephone:              "09100001111",
			Email:                  "example.com",
			PersonInChargeName:     "person",
			PersonInChargeNameKana: toOptString(&optString),
			Address: ogenlib.Address{
				PostalCode: "8891111",
				PrefID:     1,
				Address1:   "宮崎市",
				Address2:   "佐土原",
			},
		},
	}
	expectErrRes := &ogenlib.PostCreateCustomerInternalServerError{
		AccessControlAllowHeaders: toOptString(&as),
		AccessControlAllowMethods: toOptString(&as),
		AccessControlAllowOrigin:  toOptString(&as),
		Response: ogenlib.ErrorModel{
			Type:    string(ogen.InternalServerError),
			Message: "unexpected error has occurred.",
		},
	}
	type fields struct {
		customerAPI apiport.CustomerAPIPortInterface
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
		{
			name: "adapter/ogen - PostCreateCustomer - success",
			fields: fields{
				customerAPI: portMock,
			},
			args: args{
				ctx: ctx,
				req: &req,
			},
			want:      expect,
			assertion: assert.NoError,
		},
		{
			name: "adapter/ogen - PostCreateCustomer - error",
			fields: fields{
				customerAPI: portMock2,
			},
			args: args{
				ctx: ctx,
				req: &req,
			},
			want:      expectErrRes,
			assertion: assert.NoError,
		},
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
	optString := "テスト名"
	params := &ogenlib.PostSearchCustomerRequest{
		Conditions: ogenlib.PostSearchCustomerRequestConditions{},
		Pagination: ogenlib.Pagination{
			Size:   100,
			Number: 1,
		},
	}

	portModel := &apiport.Customer{
		ID:                     11,
		Name:                   "testname",
		NameKana:               &optString,
		Telephone:              "09100001111",
		Email:                  "example.com",
		PersonInChargeName:     "person",
		PersonInChargeNameKana: &optString,
		PostalCode:             "8891111",
		PrefID:                 1,
		Address1:               "宮崎市",
		Address2:               "佐土原",
	}
	pageResult := apiport.PageResult{
		Size:    int64(100),
		Total:   int64(121),
		Current: int64(1),
	}
	portResult := &apiport.CustomerSearchResult{
		CustomerList: []*apiport.Customer{portModel},
		Page:         pageResult,
	}
	as := "*"
	expect := &ogenlib.PostSearchCustomer200ResponseHeaders{
		AccessControlAllowHeaders: toOptString(&as),
		AccessControlAllowMethods: toOptString(&as),
		AccessControlAllowOrigin:  toOptString(&as),
		Response: ogenlib.PostSearchCustomer200Response{
			Page: ogenlib.Page{
				Size:    100,
				Current: 1,
				Total:   121,
			},
			Customers: []ogenlib.Customer{
				{
					ID:                     11,
					Name:                   "testname",
					NameKana:               toOptString(&optString),
					Telephone:              "09100001111",
					Email:                  "example.com",
					PersonInChargeName:     "person",
					PersonInChargeNameKana: toOptString(&optString),
					Address: ogenlib.Address{
						PostalCode: "8891111",
						PrefID:     1,
						Address1:   "宮崎市",
						Address2:   "佐土原",
					},
				},
			},
		},
	}
	expectErrRes := &ogenlib.PostSearchCustomerInternalServerError{
		AccessControlAllowHeaders: toOptString(&as),
		AccessControlAllowMethods: toOptString(&as),
		AccessControlAllowOrigin:  toOptString(&as),
		Response: ogenlib.ErrorModel{
			Type:    string(ogen.InternalServerError),
			Message: "unexpected error has occurred.",
		},
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	portMock := apiportMock.NewMockCustomerAPIPortInterface(ctrl)
	portMock.EXPECT().SearchCustomer(
		params.Pagination.Number, params.Pagination.Size, &apiport.SearchConditions{},
	).Return(portResult, nil)
	mockErr := apiport.NewAPIUnexpectedError(errors.New("test error"))
	portMock2 := apiportMock.NewMockCustomerAPIPortInterface(ctrl)
	portMock2.EXPECT().SearchCustomer(
		params.Pagination.Number, params.Pagination.Size, &apiport.SearchConditions{},
	).Return(nil, mockErr)
	ctx := context.Background()
	type fields struct {
		customerAPI apiport.CustomerAPIPortInterface
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
		{
			name: "adapter/ogen - PostSearchCustomer - success",
			fields: fields{
				customerAPI: portMock,
			},
			args: args{
				ctx: ctx,
				req: params,
			},
			want:      expect,
			assertion: assert.NoError,
		},
		{
			name: "adapter/ogen - PostSearchCustomer - error",
			fields: fields{
				customerAPI: portMock2,
			},
			args: args{
				ctx: ctx,
				req: params,
			},
			want:      expectErrRes,
			assertion: assert.NoError,
		},
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
	optString := "テスト名"
	req := ogenlib.PutModifyCustomerByIDRequest{
		ID:                     11,
		Name:                   "testname",
		NameKana:               toOptString(&optString),
		Telephone:              "09100001111",
		Email:                  "example.com",
		PersonInChargeName:     "person",
		PersonInChargeNameKana: toOptString(&optString),
		Address: ogenlib.Address{
			PostalCode: "8891111",
			PrefID:     1,
			Address1:   "宮崎市",
			Address2:   "佐土原",
		},
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
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	portMock := apiportMock.NewMockCustomerAPIPortInterface(ctrl)
	portMock.EXPECT().UpdateByID(portModel).Return(portModel, nil)
	portMock2 := apiportMock.NewMockCustomerAPIPortInterface(ctrl)
	portErr := apiport.NewAPICustomerNotFoundError(
		errors.New("error message"),
		apiport.CustomerID(11),
	)
	portMock2.EXPECT().UpdateByID(portModel).Return(nil, portErr)
	ctx := context.Background()
	as := "*"
	expect := &ogenlib.CustomerHeaders{
		AccessControlAllowHeaders: toOptString(&as),
		AccessControlAllowMethods: toOptString(&as),
		AccessControlAllowOrigin:  toOptString(&as),
		Response: ogenlib.Customer{
			ID:                     11,
			Name:                   "testname",
			NameKana:               toOptString(&optString),
			Telephone:              "09100001111",
			Email:                  "example.com",
			PersonInChargeName:     "person",
			PersonInChargeNameKana: toOptString(&optString),
			Address: ogenlib.Address{
				PostalCode: "8891111",
				PrefID:     1,
				Address1:   "宮崎市",
				Address2:   "佐土原",
			},
		},
	}
	expectErrRes := &ogenlib.PutModifyCustomerByIDNotFound{
		AccessControlAllowHeaders: toOptString(&as),
		AccessControlAllowMethods: toOptString(&as),
		AccessControlAllowOrigin:  toOptString(&as),
		Response: ogenlib.ErrorModel{
			Type:    string(ogen.ResourceNotFound),
			Message: "customer not found.",
		},
	}
	type fields struct {
		customerAPI apiport.CustomerAPIPortInterface
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
		{
			name: "adapter/ogen - PutModifyCustomerByID - success",
			fields: fields{
				customerAPI: portMock,
			},
			args: args{
				ctx: ctx,
				req: &req,
				params: ogenlib.PutModifyCustomerByIDParams{
					CustomerID: 11,
				},
			},
			want:      expect,
			assertion: assert.NoError,
		},
		{
			name: "adapter/ogen - PutModifyCustomerByID - error",
			fields: fields{
				customerAPI: portMock2,
			},
			args: args{
				ctx: ctx,
				req: &req,
				params: ogenlib.PutModifyCustomerByIDParams{
					CustomerID: 11,
				},
			},
			want:      expectErrRes,
			assertion: assert.NoError,
		},
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
	optString := "テスト名"
	params := ogenlib.GetCustomerByIDParams{
		CustomerID: 11,
	}

	portModel := &apiport.Customer{
		ID:                     11,
		Name:                   "testname",
		NameKana:               &optString,
		Telephone:              "09100001111",
		Email:                  "example.com",
		PersonInChargeName:     "person",
		PersonInChargeNameKana: &optString,
		PostalCode:             "8891111",
		PrefID:                 1,
		Address1:               "宮崎市",
		Address2:               "佐土原",
	}
	as := "*"
	expect := &ogenlib.CustomerHeaders{
		AccessControlAllowHeaders: toOptString(&as),
		AccessControlAllowMethods: toOptString(&as),
		AccessControlAllowOrigin:  toOptString(&as),
		Response: ogenlib.Customer{
			ID:                     11,
			Name:                   "testname",
			NameKana:               toOptString(&optString),
			Telephone:              "09100001111",
			Email:                  "example.com",
			PersonInChargeName:     "person",
			PersonInChargeNameKana: toOptString(&optString),
			Address: ogenlib.Address{
				PostalCode: "8891111",
				PrefID:     1,
				Address1:   "宮崎市",
				Address2:   "佐土原",
			},
		},
	}
	expectErrRes := &ogenlib.GetCustomerByIDNotFound{
		AccessControlAllowHeaders: toOptString(&as),
		AccessControlAllowMethods: toOptString(&as),
		AccessControlAllowOrigin:  toOptString(&as),
		Response: ogenlib.ErrorModel{
			Type:    string(ogen.ResourceNotFound),
			Message: "customer not found.",
		},
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	portMock := apiportMock.NewMockCustomerAPIPortInterface(ctrl)
	portMock.EXPECT().GetByID(apiport.CustomerID(11)).Return(portModel, nil)
	portMock2 := apiportMock.NewMockCustomerAPIPortInterface(ctrl)
	portErr := apiport.NewAPICustomerNotFoundError(
		errors.New("error message"),
		apiport.CustomerID(11),
	)
	portMock2.EXPECT().GetByID(apiport.CustomerID(11)).Return(nil, portErr)
	ctx := context.Background()
	type fields struct {
		customerAPI apiport.CustomerAPIPortInterface
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
		{
			name: "adapter/ogen - GetCustomerByID - success",
			fields: fields{
				customerAPI: portMock,
			},
			args: args{
				ctx:    ctx,
				params: params,
			},
			want:      expect,
			assertion: assert.NoError,
		},
		{
			name: "adapter/ogen - GetCustomerByID - error",
			fields: fields{
				customerAPI: portMock2,
			},
			args: args{
				ctx:    ctx,
				params: params,
			},
			want:      expectErrRes,
			assertion: assert.NoError,
		},
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

func TestNewEnAPIAdapter(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	portMock := apiportMock.NewMockCustomerAPIPortInterface(ctrl)
	type args struct {
		customerAPI apiport.CustomerAPIPortInterface
	}
	tests := []struct {
		name string
		args args
		want *ogen.EnAPIAdapter
	}{
		{
			name: "adapter/ogen - NewEnAPIAdapter - success",
			args: args{
				customerAPI: portMock,
			},
			want: &ogen.EnAPIAdapter{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(
				t,
				reflect.TypeOf(tt.want),
				reflect.TypeOf(ogen.NewEnAPIAdapter(tt.args.customerAPI)))
		})
	}
}
