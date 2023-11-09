package rdbadapter

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"

	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
)

func (rdb *MySQL) GetCustomerByID(customerID int64) (*Customer, error) {
	db := bun.NewDB(rdb.Conn, mysqldialect.New())
	customer := Customer{}
	getErr := db.NewSelect().Model(&customer).Where("? = ?", bun.Ident("id"), customerID).Scan(context.Background())
	if getErr != nil {
		if errors.Is(getErr, sql.ErrNoRows) {
			return nil, ErrRdbCustomerNotFound
		}
		slog.Error(fmt.Sprintf("db error log. %v", getErr))
		return nil, ErrRdbUnexpected
	}
	return &customer, nil
}
