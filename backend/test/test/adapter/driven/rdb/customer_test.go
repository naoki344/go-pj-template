package rdbadapter_test

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/cockroachdb/errors"
	rdbadapter "github.com/g-stayfresh/en/backend/internal/adapter/driven/rdb"
	"github.com/stretchr/testify/assert"
)

var (
	columns                = []string{"id", "name", "name_kana", "telephone", "email", "person_in_charge_name", "person_in_charge_name_kana", "postal_code", "pref_id", "address1", "address2"} //nolint:gochecknoglobals
	nameKana               = "テストメイ"                                                                                                                                                            //nolint:gochecknoglobals
	personInChargeNameKana = "タントウシャA"                                                                                                                                                          //nolint:gochecknoglobals
	//nolint:gochecknoglobals
	customerData = rdbadapter.Customer{
		ID:                     11,
		Name:                   "テスト名",
		NameKana:               &nameKana,
		Telephone:              "09011112222",
		Email:                  "example@example.comm",
		PersonInChargeName:     "担当者A",
		PersonInChargeNameKana: &personInChargeNameKana,
		PostalCode:             "8801111",
		PrefID:                 1,
		Address1:               "宮崎市佐土原町上田島",
		Address2:               "111-1111",
	}
)

//nolint:gochecknoglobals
var newCustomerData = rdbadapter.Customer{
	Name:                   "テスト名",
	NameKana:               &nameKana,
	Telephone:              "09011112222",
	Email:                  "example@example.comm",
	PersonInChargeName:     "担当者A",
	PersonInChargeNameKana: &personInChargeNameKana,
	PostalCode:             "8801111",
	PrefID:                 1,
	Address1:               "宮崎市佐土原町上田島",
	Address2:               "111-1111",
}

func TestMySQL_GetCustomerByIDFull(t *testing.T) {
	expectQuery := "SELECT `customer`.`id`, `customer`.`name`, `customer`.`name_kana`, `customer`.`telephone`, `customer`.`email`, `customer`.`person_in_charge_name`, `customer`.`person_in_charge_name_kana`, `customer`.`postal_code`, `customer`.`pref_id`, `customer`.`address1`, `customer`.`address2` FROM `customers` AS `customer` WHERE \\(`id` = 11\\)"
	db, mock, err := sqlmock.New()
	mock.ExpectQuery(expectQuery).
		WillReturnRows(sqlmock.NewRows(columns).AddRow(
			customerData.ID,
			customerData.Name,
			*customerData.NameKana,
			customerData.Telephone,
			customerData.Email,
			customerData.PersonInChargeName,
			*customerData.PersonInChargeNameKana,
			customerData.PostalCode,
			customerData.PrefID,
			customerData.Address1,
			customerData.Address2))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	db2, mock2, err := sqlmock.New()
	mock2.ExpectQuery(expectQuery).
		WillReturnError(sql.ErrNoRows)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db2.Close()
	errMock := errors.New("unexpected error")
	db3, mock3, err := sqlmock.New()
	mock3.ExpectQuery(expectQuery).
		WillReturnError(errMock)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db3.Close()
	testCustomerID := int64(11)
	type fields struct {
		Conn *sql.DB
	}
	type args struct {
		customerID int64
	}

	var notFoundError *rdbadapter.RdbCustomerNotFoundError
	var unexpectedError *rdbadapter.RdbUnexpectedError

	tests := []struct {
		name      string
		fields    fields
		args      args
		want      *rdbadapter.Customer
		assertion assert.ErrorAssertionFunc
	}{
		{
			name:   "adapter/rdb GetCustomerByID Test - success",
			fields: fields{db},
			args: args{
				customerID: testCustomerID,
			},
			want:      &customerData,
			assertion: assert.NoError,
		},
		{
			name:   "adapter/rdb GetCustomerByID Test - error(NoRows)",
			fields: fields{db2},
			args: args{
				customerID: testCustomerID,
			},
			want: nil,
			assertion: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorAs(t, err, &notFoundError)
			},
		},
		{
			name:   "adapter/rdb GetCustomerByID Test - error(other)",
			fields: fields{db3},
			args: args{
				customerID: testCustomerID,
			},
			want: nil,
			assertion: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorAs(t, err, &unexpectedError)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rdb := &rdbadapter.MySQL{
				Conn: tt.fields.Conn,
			}
			got, err := rdb.GetCustomerByID(tt.args.customerID)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMySQL_UpdateCustomerByID(t *testing.T) {
	expectQuery := "UPDATE `customers` AS `customer` SET" +
		fmt.Sprintf(" `name` = '%s'", customerData.Name) +
		fmt.Sprintf(", `name_kana` = '%s'", *customerData.NameKana) +
		fmt.Sprintf(", `telephone` = '%s'", customerData.Telephone) +
		fmt.Sprintf(", `email` = '%s'", customerData.Email) +
		fmt.Sprintf(", `person_in_charge_name` = '%s'", customerData.PersonInChargeName) +
		fmt.Sprintf(", `person_in_charge_name_kana` = '%s'", *customerData.PersonInChargeNameKana) +
		fmt.Sprintf(", `postal_code` = '%s'", customerData.PostalCode) +
		fmt.Sprintf(", `pref_id` = %d", customerData.PrefID) +
		fmt.Sprintf(", `address1` = '%s'", customerData.Address1) +
		fmt.Sprintf(", `address2` = '%s'", customerData.Address2)
	db, mock, err := sqlmock.New()
	mock.ExpectExec(expectQuery).
		WillReturnResult(sqlmock.NewResult(1010, 1))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	db2, mock2, err2 := sqlmock.New()
	mock2.ExpectExec(expectQuery).
		WillReturnError(sql.ErrNoRows)
	if err2 != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db2.Close()
	var notFoundError *rdbadapter.RdbCustomerNotFoundError

	type fields struct {
		Conn *sql.DB
	}
	type args struct {
		customer *rdbadapter.Customer
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		assertion assert.ErrorAssertionFunc
	}{
		{
			name:   "adapter/rdb UpdateCustomerByID Test - success",
			fields: fields{db},
			args: args{
				customer: &customerData,
			},
			assertion: assert.NoError,
		},
		{
			name:   "adapter/rdb UpdateCustomerByID Test - error(noRows)",
			fields: fields{db2},
			args: args{
				customer: &customerData,
			},
			assertion: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorAs(t, err, &notFoundError)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rdb := &rdbadapter.MySQL{
				Conn: tt.fields.Conn,
			}
			tt.assertion(t, rdb.UpdateCustomerByID(tt.args.customer))
		})
	}
}

func TestMySQL_InsertCustomer(t *testing.T) {
	expectQuery := "INSERT INTO `customers` \\(`id`, `name`, `name_kana`, `telephone`, `email`, `person_in_charge_name`, `person_in_charge_name_kana`, `postal_code`, `pref_id`, `address1`, `address2`\\) " +
		"VALUES \\(" +
		"DEFAULT" +
		fmt.Sprintf(", '%s'", newCustomerData.Name) +
		fmt.Sprintf(", '%s'", *newCustomerData.NameKana) +
		fmt.Sprintf(", '%s'", newCustomerData.Telephone) +
		fmt.Sprintf(", '%s'", newCustomerData.Email) +
		fmt.Sprintf(", '%s'", newCustomerData.PersonInChargeName) +
		fmt.Sprintf(", '%s'", *newCustomerData.PersonInChargeNameKana) +
		fmt.Sprintf(", '%s'", newCustomerData.PostalCode) +
		fmt.Sprintf(", %d", newCustomerData.PrefID) +
		fmt.Sprintf(", '%s'", newCustomerData.Address1) +
		fmt.Sprintf(", '%s'", newCustomerData.Address2) +
		"\\)"
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mock.ExpectExec(expectQuery).
		WillReturnResult(sqlmock.NewResult(1010, 1))
	newCustomerExpect := newCustomerData
	newCustomerExpect.ID = 1010
	type fields struct {
		Conn *sql.DB
	}
	type args struct {
		customer *rdbadapter.Customer
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      *rdbadapter.Customer
		assertion assert.ErrorAssertionFunc
	}{
		{
			name:   "adapter/rdb UpdateCustomerByID Test - success",
			fields: fields{db},
			args: args{
				customer: &newCustomerData,
			},
			want:      &newCustomerExpect,
			assertion: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rdb := &rdbadapter.MySQL{
				Conn: tt.fields.Conn,
			}
			got, err := rdb.InsertCustomer(tt.args.customer)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMySQL_SearchCustomer(t *testing.T) {
	expectQuery := "SELECT `customer`.`id`, `customer`.`name`, `customer`.`name_kana`, `customer`.`telephone`, `customer`.`email`, `customer`.`person_in_charge_name`, `customer`.`person_in_charge_name_kana`, `customer`.`postal_code`, `customer`.`pref_id`, `customer`.`address1`, `customer`.`address2` FROM `customers` AS `customer` WHERE \\(`id` between 901 and 1000\\)"
	db, mock, err := sqlmock.New()
	mock.ExpectQuery(expectQuery).
		WillReturnRows(sqlmock.NewRows(columns).AddRow(
			customerData.ID,
			customerData.Name,
			*customerData.NameKana,
			customerData.Telephone,
			customerData.Email,
			customerData.PersonInChargeName,
			*customerData.PersonInChargeNameKana,
			customerData.PostalCode,
			customerData.PrefID,
			customerData.Address1,
			customerData.Address2))
	mock.ExpectQuery("SELECT").
		WillReturnRows(sqlmock.NewRows([]string{"Total"}).AddRow(101))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	db2, mock2, err := sqlmock.New()
	mock2.ExpectQuery(expectQuery).
		WillReturnRows(sqlmock.NewRows(columns))
	mock2.ExpectQuery("SELECT").
		WillReturnRows(sqlmock.NewRows([]string{"Total"}).AddRow(0))

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db2.Close()
	type fields struct {
		Conn *sql.DB
	}

	type args struct {
		pageNumber int64
		pageSize   int64
		conditions *rdbadapter.SearchConditions
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      *rdbadapter.CustomerSearchResult
		assertion assert.ErrorAssertionFunc
	}{
		{
			name:   "adapter/rdb UpdateCustomerByID Test - success",
			fields: fields{db},
			args: args{
				pageNumber: 10,
				pageSize:   100,
				conditions: &rdbadapter.SearchConditions{},
			},
			want: &rdbadapter.CustomerSearchResult{
				PageInfo: rdbadapter.PageInfo{
					Current: 10,
					Size:    100,
					Total:   101,
				},
				CustomerList: []rdbadapter.Customer{customerData},
			},
			assertion: assert.NoError,
		},
		{
			name:   "adapter/rdb UpdateCustomerByID Test - success(no row)",
			fields: fields{db2},
			args: args{
				pageNumber: 10,
				pageSize:   100,
				conditions: &rdbadapter.SearchConditions{},
			},
			want: &rdbadapter.CustomerSearchResult{
				PageInfo: rdbadapter.PageInfo{
					Current: 10,
					Size:    100,
					Total:   0,
				},
				CustomerList: rdbadapter.CustomerList{},
			},
			assertion: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rdb := &rdbadapter.MySQL{
				Conn: tt.fields.Conn,
			}
			got, err := rdb.SearchCustomer(tt.args.pageNumber, tt.args.pageSize, tt.args.conditions)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
