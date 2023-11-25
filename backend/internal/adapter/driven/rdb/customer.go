package rdbadapter

import (
	"context"
	"database/sql"
	"log/slog"

	"github.com/cockroachdb/errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"github.com/uptrace/bun/extra/bundebug"
)

const CustomersTableName = "customers"

func (rdb *MySQL) GetCustomerByID(customerID int64) (*Customer, error) {
	db := bun.NewDB(rdb.Conn, mysqldialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	customer := Customer{}
	getErr := db.NewSelect().Model(&customer).Where("? = ?", bun.Ident("id"), customerID).Scan(context.Background())
	if getErr != nil {
		slog.Error("db error log.", "error", getErr)
		if errors.Is(getErr, sql.ErrNoRows) {
			return nil, NewRdbCustomerNotFoundError(getErr, customerID)
		}
		return nil, NewRdbUnexpectedError(getErr)
	}
	slog.Info("customer get success.")
	return &customer, nil
}

func (rdb *MySQL) UpdateCustomerByID(customer *Customer) error {
	db := bun.NewDB(rdb.Conn, mysqldialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	res, err := db.NewUpdate().Model(customer).WherePK().Exec(context.Background())
	if err != nil {
		slog.Error("db error log.", "error", err)
		if errors.Is(err, sql.ErrNoRows) {
			return NewRdbCustomerNotFoundError(err, customer.ID)
		}
		return NewRdbUnexpectedError(err)
	}
	slog.Info("customer update success.", slog.Any("result", &res))
	return nil
}

func (rdb *MySQL) InsertCustomer(customer *Customer) (*Customer, error) {
	db := bun.NewDB(rdb.Conn, mysqldialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	res, err := db.NewInsert().Model(customer).Exec(context.Background())
	if err != nil {
		slog.Error("db error log.", "error", err)
		return nil, NewRdbUnexpectedError(err)
	}
	slog.Info("customer update success.", slog.Any("result", &res))
	return customer, nil
}

func (rdb *MySQL) SearchCustomer(pageNumber int64, pageSize int64, conditions *SearchConditions) (*CustomerSearchResult, error) {
	db := bun.NewDB(rdb.Conn, mysqldialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	customers := CustomerList{}
	start, end := calcPaging(pageNumber, pageSize)
	getErr := db.NewSelect().Model(&customers).Where(
		"? between ? and ?", bun.Ident("id"), start, end).Scan(context.Background())
	if getErr != nil {
		slog.Error("db error log.", "error", getErr)
		return nil, NewRdbUnexpectedError(getErr)
	}

	var total int64
	totalCountErr := db.NewSelect().
		ColumnExpr("COUNT(id) AS Total").
		TableExpr(CustomersTableName).
		Scan(context.Background(), &total)
	if totalCountErr != nil {
		slog.Error("db error log.", "error", totalCountErr)
		return nil, NewRdbUnexpectedError(totalCountErr)
	}
	slog.Info("customer search success.")
	result := &CustomerSearchResult{
		CustomerList: customers,
		PageInfo: PageInfo{
			Size:    pageSize,
			Total:   total,
			Current: pageNumber,
		},
	}
	return result, nil
}
