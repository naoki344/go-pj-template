package rdbport_test

import (
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	rdbadapter "github.com/naoki344/go-pj-template/backend/internal/adapter/driven/rdb"
	customermodel "github.com/naoki344/go-pj-template/backend/internal/domain/model/customer"
	pagemodel "github.com/naoki344/go-pj-template/backend/internal/domain/model/page"
	rdbport "github.com/naoki344/go-pj-template/backend/internal/port/driven/rdb"
	adaptermock "github.com/naoki344/go-pj-template/backend/test/mock/adapter/driven/rdb"
	"github.com/stretchr/testify/assert"
)

func TestNewRdbPort(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mock := adaptermock.NewMockRdbInterface(ctrl)
	type args struct {
		rdb rdbadapter.RdbInterface
	}
	tests := []struct {
		name string
		args args
		want *rdbport.RdbPort
	}{
		{
			name: "port/rdb - NewRdbPort - success",
			args: args{
				rdb: mock,
			},
			want: &rdbport.RdbPort{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, reflect.TypeOf(tt.want), reflect.TypeOf(rdbport.NewRdbPort(tt.args.rdb)))
		})
	}
}

func TestRdbPort_CustomerCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
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
		ID:                     customermodel.ID(111),
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
	newCustomerData := &rdbadapter.Customer{
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
	createdCustomerData := &rdbadapter.Customer{
		ID:                     111,
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
	mock := adaptermock.NewMockRdbInterface(ctrl)
	mock.EXPECT().InsertCustomer(newCustomerData).Return(createdCustomerData, nil)
	type fields struct {
		rdb rdbadapter.RdbInterface
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
			name: "port/rdb - CustomerCreate - success",
			fields: fields{
				rdb: mock,
			},
			args: args{
				customer: customerReqModel,
			},
			want:      customerResModel,
			assertion: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			port := rdbport.NewRdbPort(tt.fields.rdb)
			got, err := port.CustomerCreate(tt.args.customer)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestRdbPort_CustomerGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	optString := "テスト名"
	customerResModel := &customermodel.Customer{
		ID:                     customermodel.ID(111),
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
	customerData := &rdbadapter.Customer{
		ID:                     111,
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
	mock := adaptermock.NewMockRdbInterface(ctrl)
	mock.EXPECT().GetCustomerByID(int64(111)).Return(customerData, nil)
	type fields struct {
		rdb rdbadapter.RdbInterface
	}
	type args struct {
		customerID customermodel.ID
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      *customermodel.Customer
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "port/rdb - CustomerGet - success",
			fields: fields{
				rdb: mock,
			},
			args: args{
				customerID: customermodel.ID(111),
			},
			want:      customerResModel,
			assertion: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			port := rdbport.NewRdbPort(tt.fields.rdb)
			got, err := port.CustomerGet(tt.args.customerID)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestRdbPort_CustomerUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	optString := "テスト名"
	customerModel := &customermodel.Customer{
		ID:                     customermodel.ID(111),
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
			Address2:   customermodel.Address2("佐土原aaaaaaaaaaa"),
		},
	}
	newCustomerData := &rdbadapter.Customer{
		ID:                     111,
		Name:                   "testname",
		NameKana:               &optString,
		Telephone:              "09100001111",
		Email:                  "example.com",
		PersonInChargeName:     "person",
		PersonInChargeNameKana: &optString,
		PostalCode:             "8891111",
		PrefID:                 1,
		Address1:               "宮崎市",
		Address2:               "佐土原aaaaaaaaaaa",
	}
	mock := adaptermock.NewMockRdbInterface(ctrl)
	mock.EXPECT().UpdateCustomerByID(newCustomerData).Return(nil)
	type fields struct {
		rdb rdbadapter.RdbInterface
	}
	type args struct {
		customer *customermodel.Customer
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "port/rdb - CustomerGet - success",
			fields: fields{
				rdb: mock,
			},
			args: args{
				customer: customerModel,
			},
			assertion: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			port := rdbport.NewRdbPort(tt.fields.rdb)
			tt.assertion(t, port.CustomerUpdate(tt.args.customer))
		})
	}
}

func TestRdbPort_CustomerSearch(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	optString := "テスト名"
	customerModel := &customermodel.Customer{
		ID:                     customermodel.ID(111),
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
			Address2:   customermodel.Address2("佐土原aaaaaaaaaaa"),
		},
	}
	customerData := &rdbadapter.Customer{
		ID:                     111,
		Name:                   "testname",
		NameKana:               &optString,
		Telephone:              "09100001111",
		Email:                  "example.com",
		PersonInChargeName:     "person",
		PersonInChargeNameKana: &optString,
		PostalCode:             "8891111",
		PrefID:                 1,
		Address1:               "宮崎市",
		Address2:               "佐土原aaaaaaaaaaa",
	}
	mock := adaptermock.NewMockRdbInterface(ctrl)
	adapterResult := &rdbadapter.CustomerSearchResult{
		PageInfo: rdbadapter.PageInfo{
			Size:    int64(100),
			Current: int64(11),
			Total:   int64(11110),
		},
		CustomerList: []rdbadapter.Customer{
			*customerData,
		},
	}

	mock.EXPECT().SearchCustomer(int64(11), int64(100), &rdbadapter.SearchConditions{}).Return(
		adapterResult, nil)
	type fields struct {
		rdb rdbadapter.RdbInterface
	}
	type args struct {
		pageNumber int64
		pageSize   int64
		conditions *customermodel.SearchConditions
	}
	result := &[]*customermodel.Customer{
		customerModel,
	}
	result1 := &pagemodel.PageResult{
		Size:    pagemodel.Size(100),
		Current: pagemodel.Current(11),
		Total:   pagemodel.Total(11110),
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
			name: "port/rdb - CustomerGet - success",
			fields: fields{
				rdb: mock,
			},
			args: args{
				pageNumber: int64(11),
				pageSize:   int64(100),
				conditions: &customermodel.SearchConditions{},
			},
			want:      result,
			want1:     result1,
			assertion: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			port := rdbport.NewRdbPort(tt.fields.rdb)
			got, got1, err := port.CustomerSearch(tt.args.pageNumber, tt.args.pageSize, tt.args.conditions)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}
