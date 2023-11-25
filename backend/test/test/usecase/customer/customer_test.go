package customerusecase_test

import (
	"reflect"
	"testing"

	"github.com/cockroachdb/errors"
	customermodel "github.com/g-stayfresh/en/backend/internal/domain/model/customer"
	pagemodel "github.com/g-stayfresh/en/backend/internal/domain/model/page"
	rdbport "github.com/g-stayfresh/en/backend/internal/port/driven/rdb"
	testTarget "github.com/g-stayfresh/en/backend/internal/usecase/customer"
	rdbportMock "github.com/g-stayfresh/en/backend/test/mock/port/driven/rdb"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestNewCustomerUsecase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		port rdbport.RdbPortInterface
	}
	tests := []struct {
		name string
		args args
		want *testTarget.CustomerUsecase
	}{
		{
			name: "New CustomerUsecase",
			args: args{
				port: rdbportMock.NewMockRdbPortInterface(ctrl),
			},
			want: &testTarget.CustomerUsecase{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, reflect.TypeOf(tt.want), reflect.TypeOf(testTarget.NewCustomerUsecase(tt.args.port)))
		})
	}
}

func TestCustomerUsecase_GetByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		port rdbport.RdbPortInterface
	}
	type args struct {
		customerID customermodel.ID
	}

	mockInput := customermodel.ID(11)
	mockExpect := customermodel.Customer{}
	mock := rdbportMock.NewMockRdbPortInterface(ctrl)
	mock.EXPECT().CustomerGet(mockInput).Return(&mockExpect, nil)
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      *customermodel.Customer
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "usecase/customer - GetByID - success",
			fields: fields{
				port: mock,
			},
			args: args{
				customerID: mockInput,
			},
			want:      &mockExpect,
			assertion: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// func (usecase *CustomerUsecase) GetByID(customerID customermodel.ID) (*customermodel.Customer, error) {
			usecase := testTarget.NewCustomerUsecase(tt.fields.port)
			got, err := usecase.GetByID(tt.args.customerID)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCustomerUsecase_UpdateByID(t *testing.T) {
	customer := customermodel.Customer{ID: 1}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mock := rdbportMock.NewMockRdbPortInterface(ctrl)
	mock.EXPECT().CustomerUpdate(&customer).Return(nil)
	mock.EXPECT().CustomerGet(customer.ID).Return(&customer, nil)

	mock2 := rdbportMock.NewMockRdbPortInterface(ctrl)
	mock2.EXPECT().CustomerUpdate(&customer).Return(nil)
	mock2.EXPECT().CustomerGet(customer.ID).Return(nil, errors.New("error"))

	type fields struct {
		port rdbport.RdbPortInterface
	}
	type args struct {
		customer *customermodel.Customer
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      *customermodel.Customer
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "usecase/customer - UpdateByID - success",
			fields: fields{
				port: mock,
			},
			args: args{
				customer: &customer,
			},
			want:      &customer,
			assertion: assert.NoError,
		},
		{
			name: "usecase/customer - UpdateByID - failure",
			fields: fields{
				port: mock2,
			},
			args: args{
				customer: &customer,
			},
			want:      nil,
			assertion: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			usecase := testTarget.NewCustomerUsecase(tt.fields.port)
			got, err := usecase.UpdateByID(tt.args.customer)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCustomerUsecase_Create(t *testing.T) {
	customer := customermodel.Customer{}
	createdCustomer := customermodel.Customer{ID: 11}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mock := rdbportMock.NewMockRdbPortInterface(ctrl)
	mock.EXPECT().CustomerCreate(&customer).Return(&createdCustomer, nil)
	type fields struct {
		port rdbport.RdbPortInterface
	}
	type args struct {
		customer *customermodel.Customer
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      *customermodel.Customer
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "usecase/customer - PostByID - success",
			fields: fields{
				port: mock,
			},
			args: args{
				customer: &customer,
			},
			want:      &createdCustomer,
			assertion: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			usecase := testTarget.NewCustomerUsecase(tt.fields.port)
			got, err := usecase.Create(tt.args.customer)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCustomerUsecase_Search(t *testing.T) {
	customer := customermodel.Customer{}
	pageRes := pagemodel.PageResult{}
	searchCustomers := []*customermodel.Customer{&customer}
	searchConditions := &customermodel.SearchConditions{}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mock := rdbportMock.NewMockRdbPortInterface(ctrl)
	pageNumber := int64(1)
	pageSize := int64(100)
	mock.EXPECT().CustomerSearch(pageNumber, pageSize, searchConditions).Return(&searchCustomers, &pageRes, nil)
	type fields struct {
		port rdbport.RdbPortInterface
	}
	type args struct {
		pageNumber int64
		pageSize   int64
		conditions *customermodel.SearchConditions
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      *[]*customermodel.Customer
		want1     *pagemodel.PageResult
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "usecase/customer - PostSearch - success",
			fields: fields{
				port: mock,
			},
			args: args{
				pageNumber: pageNumber,
				pageSize:   pageSize,
				conditions: searchConditions,
			},
			want:      &searchCustomers,
			want1:     &pageRes,
			assertion: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			usecase := testTarget.NewCustomerUsecase(tt.fields.port)
			got, got1, err := usecase.Search(tt.args.pageNumber, tt.args.pageSize, tt.args.conditions)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}
