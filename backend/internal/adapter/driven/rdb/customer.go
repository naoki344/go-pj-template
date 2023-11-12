package rdbadapter

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"

	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"github.com/uptrace/bun/extra/bundebug"
)

func (rdb *MySQL) GetCustomerByID(customerID int64) (*Customer, error) {
	db := bun.NewDB(rdb.Conn, mysqldialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	customer := Customer{}
	getErr := db.NewSelect().Model(&customer).Where("? = ?", bun.Ident("id"), customerID).Scan(context.Background())
	if getErr != nil {
		slog.Error("db error log.", "error", getErr)
		if errors.Is(getErr, sql.ErrNoRows) {
			return nil, ErrRdbCustomerNotFound
		}
		return nil, ErrRdbUnexpected
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
			return ErrRdbCustomerNotFound
		}
		return ErrRdbUnexpected
	}
	slog.Info("customer update success.", slog.Any("result", &res))
	return nil
}
