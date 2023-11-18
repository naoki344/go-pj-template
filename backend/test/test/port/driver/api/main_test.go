package apiport_test

import (
	"errors"
	"reflect"
	"testing"

	errormodel "github.com/g-stayfresh/en/backend/internal/domain/error"
	customermodel "github.com/g-stayfresh/en/backend/internal/domain/model/customer"
	pagemodel "github.com/g-stayfresh/en/backend/internal/domain/model/page"
	apiport "github.com/g-stayfresh/en/backend/internal/port/driver/api"
	usecase "github.com/g-stayfresh/en/backend/internal/usecase/customer"
	usecaseMock "github.com/g-stayfresh/en/backend/test/mock/usecase/customer"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestNewCustomerAPIPort(t *testing.T) {
	type args struct {
		usecase usecase.CustomerUsecaseInterface
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mock := usecaseMock.NewMockCustomerUsecaseInterface(ctrl)

	tests := []struct {
		name string
		args args
		want *apiport.CustomerAPIPort
	}{
		{
			name: "port/api - NewCustomerAPIPort - success",
			args: args{
				usecase: mock,
			},
			want: &apiport.CustomerAPIPort{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, reflect.TypeOf(tt.want), reflect.TypeOf(apiport.NewCustomerAPIPort(tt.args.usecase)))
		})
	}
}

func TestCustomerAPIPort_GetByID(t *testing.T) {
	optString := "テスト名"
	customerModel := &customermodel.Customer{
		ID:                     customermodel.ID(11),
		Name:                   customermodel.Name("testname"),
		NameKana:               customermodel.NameKana(&optString),
		Telephone:              customermodel.Telephone("09100001111"),
		Email:                  customermodel.Email("example.com"),
		PersonInChargeName:     customermodel.PersonInChargeName("person"),
		PersonInChargeNameKana: customermodel.PersonInChargeNameKana(&optString),
		Address: customermodel.Address{
			PostalCode: customermodel.PostalCode("8891111"),
			PrefID:     customermodel.PrefID(1),
			Address1:   customermodel.Address1("宮崎市"),
			Address2:   customermodel.Address2("佐土原"),
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
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mock := usecaseMock.NewMockCustomerUsecaseInterface(ctrl)
	mock.EXPECT().GetByID(customermodel.ID(11)).Return(customerModel, nil)
	mock2 := usecaseMock.NewMockCustomerUsecaseInterface(ctrl)
	mock2.EXPECT().GetByID(customermodel.ID(11)).Return(nil, errormodel.ErrCustomerNotFound)
	type fields struct {
		usecase usecase.CustomerUsecaseInterface
	}
	type args struct {
		customerID apiport.CustomerID
	}
	var expectError *apiport.APICustomerNotFoundError
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      *apiport.Customer
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "port/api - GetByID - success",
			fields: fields{
				usecase: mock,
			},
			args: args{
				apiport.CustomerID(11),
			},
			want:      portModel,
			assertion: assert.NoError,
		},
		{
			name: "port/api - GetByID - error",
			fields: fields{
				usecase: mock2,
			},
			args: args{
				apiport.CustomerID(11),
			},
			want: nil,
			assertion: func(t assert.TestingT, err error, i ...interface{}) bool {
				return errors.As(err, &expectError)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			port := apiport.NewCustomerAPIPort(tt.fields.usecase)
			got, err := port.GetByID(tt.args.customerID)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCustomerAPIPort_UpdateByID(t *testing.T) {
	optString := "テスト名"
	customerModel := &customermodel.Customer{
		ID:                     customermodel.ID(11),
		Name:                   customermodel.Name("testname"),
		NameKana:               customermodel.NameKana(&optString),
		Telephone:              customermodel.Telephone("09100001111"),
		Email:                  customermodel.Email("example.com"),
		PersonInChargeName:     customermodel.PersonInChargeName("person"),
		PersonInChargeNameKana: customermodel.PersonInChargeNameKana(&optString),
		Address: customermodel.Address{
			PostalCode: customermodel.PostalCode("8891111"),
			PrefID:     customermodel.PrefID(1),
			Address1:   customermodel.Address1("宮崎市"),
			Address2:   customermodel.Address2("佐土原"),
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
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mock := usecaseMock.NewMockCustomerUsecaseInterface(ctrl)
	mock.EXPECT().UpdateByID(customerModel).Return(customerModel, nil)
	mock2 := usecaseMock.NewMockCustomerUsecaseInterface(ctrl)
	mock2.EXPECT().UpdateByID(customerModel).Return(nil, errormodel.ErrCustomerNotFound)
	type fields struct {
		usecase usecase.CustomerUsecaseInterface
	}
	type args struct {
		customer *apiport.Customer
	}
	var expectError *apiport.APICustomerNotFoundError
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      *apiport.Customer
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "port/api - UpdateByID - success",
			fields: fields{
				usecase: mock,
			},
			args: args{
				customer: portModel,
			},
			want:      portModel,
			assertion: assert.NoError,
		},
		{
			name: "port/api - UpdateByID - error",
			fields: fields{
				usecase: mock2,
			},
			args: args{
				customer: portModel,
			},
			want: nil,
			assertion: func(t assert.TestingT, err error, i ...interface{}) bool {
				return errors.As(err, &expectError)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			port := apiport.NewCustomerAPIPort(tt.fields.usecase)
			got, err := port.UpdateByID(tt.args.customer)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCustomerAPIPort_CreateCustomer(t *testing.T) {
	optString := "テスト名"
	customerReqModel := &customermodel.Customer{
		Name:                   customermodel.Name("testname"),
		NameKana:               customermodel.NameKana(&optString),
		Telephone:              customermodel.Telephone("09100001111"),
		Email:                  customermodel.Email("example.com"),
		PersonInChargeName:     customermodel.PersonInChargeName("person"),
		PersonInChargeNameKana: customermodel.PersonInChargeNameKana(&optString),
		Address: customermodel.Address{
			PostalCode: customermodel.PostalCode("8891111"),
			PrefID:     customermodel.PrefID(1),
			Address1:   customermodel.Address1("宮崎市"),
			Address2:   customermodel.Address2("佐土原"),
		},
	}
	customerResModel := &customermodel.Customer{
		ID:                     customermodel.ID(11),
		Name:                   customermodel.Name("testname"),
		NameKana:               customermodel.NameKana(&optString),
		Telephone:              customermodel.Telephone("09100001111"),
		Email:                  customermodel.Email("example.com"),
		PersonInChargeName:     customermodel.PersonInChargeName("person"),
		PersonInChargeNameKana: customermodel.PersonInChargeNameKana(&optString),
		Address: customermodel.Address{
			PostalCode: customermodel.PostalCode("8891111"),
			PrefID:     customermodel.PrefID(1),
			Address1:   customermodel.Address1("宮崎市"),
			Address2:   customermodel.Address2("佐土原"),
		},
	}
	portReqModel := &apiport.Customer{
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
	portResModel := &apiport.Customer{
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
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mock := usecaseMock.NewMockCustomerUsecaseInterface(ctrl)
	mock.EXPECT().Create(customerReqModel).Return(customerResModel, nil)
	mock2 := usecaseMock.NewMockCustomerUsecaseInterface(ctrl)
	mock2.EXPECT().Create(customerReqModel).Return(nil, errormodel.ErrUnexpectedError)
	type fields struct {
		usecase usecase.CustomerUsecaseInterface
	}
	type args struct {
		customer *apiport.Customer
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      *apiport.Customer
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "port/api - Create - success",
			fields: fields{
				usecase: mock,
			},
			args: args{
				customer: portReqModel,
			},
			want:      portResModel,
			assertion: assert.NoError,
		},
		{
			name: "port/api - Create - error",
			fields: fields{
				usecase: mock2,
			},
			args: args{
				customer: portReqModel,
			},
			want: nil,
			assertion: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorIs(t, err, apiport.ErrUnexpected)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			port := apiport.NewCustomerAPIPort(tt.fields.usecase)
			got, err := port.CreateCustomer(tt.args.customer)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCustomerAPIPort_SearchCustomer(t *testing.T) {
	type fields struct {
		usecase usecase.CustomerUsecaseInterface
	}
	type args struct {
		pageNumber int64
		pageSize   int64
		conditions *apiport.SearchConditions
	}
	optString := "テスト名"
	customerModel := &[]*customermodel.Customer{
		{
			ID:                     customermodel.ID(11),
			Name:                   customermodel.Name("testname"),
			NameKana:               customermodel.NameKana(&optString),
			Telephone:              customermodel.Telephone("09100001111"),
			Email:                  customermodel.Email("example.com"),
			PersonInChargeName:     customermodel.PersonInChargeName("person"),
			PersonInChargeNameKana: customermodel.PersonInChargeNameKana(&optString),
			Address: customermodel.Address{
				PostalCode: customermodel.PostalCode("8891111"),
				PrefID:     customermodel.PrefID(1),
				Address1:   customermodel.Address1("宮崎市"),
				Address2:   customermodel.Address2("佐土原"),
			},
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
	// *[]*customermodel.Customer, *pagemodel.PageResult, error
	input := args{
		pageNumber: int64(1),
		pageSize:   int64(1),
		conditions: &apiport.SearchConditions{},
	}
	page := &pagemodel.PageResult{
		Size:    pagemodel.Size(1),
		Current: pagemodel.Current(1),
		Total:   pagemodel.Total(1),
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mock := usecaseMock.NewMockCustomerUsecaseInterface(ctrl)
	mock.EXPECT().Search(input.pageNumber, input.pageSize, &customermodel.SearchConditions{}).Return(customerModel, page, nil)
	mock2 := usecaseMock.NewMockCustomerUsecaseInterface(ctrl)
	mock2.EXPECT().Search(input.pageNumber, input.pageSize, &customermodel.SearchConditions{}).Return(nil, nil, errormodel.ErrUnexpectedError)
	expect := &apiport.CustomerSearchResult{
		Page: apiport.PageResult{
			Size:    int64(1),
			Current: int64(1),
			Total:   int64(1),
		},
		CustomerList: []*apiport.Customer{portModel},
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      *apiport.CustomerSearchResult
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "port/api - SearchCusomer - success",
			fields: fields{
				usecase: mock,
			},
			args:      input,
			want:      expect,
			assertion: assert.NoError,
		},
		{
			name: "port/api - SearchCusomer - error",
			fields: fields{
				usecase: mock2,
			},
			args: input,
			want: nil,
			assertion: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorIs(t, err, apiport.ErrUnexpected)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			port := apiport.NewCustomerAPIPort(tt.fields.usecase)
			got, err := port.SearchCustomer(tt.args.pageNumber, tt.args.pageSize, tt.args.conditions)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
