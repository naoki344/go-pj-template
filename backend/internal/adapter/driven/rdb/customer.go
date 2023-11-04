package rdbadapter

import (
	"context"
	"fmt"
	"log/slog"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
)


type Customer struct {
	ID      int64  `bun:"i"`
	Title   string `bun:"title"`
	Content string `bun:"content"`
}


func (rdb *MySQL) GetCustomerByID(customerID int64) (*Customer, error){
	db := bun.NewDB(rdb.Conn, mysqldialect.New())
	customer := Customer{}
	get_err := db.NewSelect().Model(&customer).Where("? = ?", bun.Ident("id"), customerID).Scan(context.Background())
	if get_err != nil {
		if get_err == sql.ErrNoRows {
			return nil, RdbErrCustomerNotFound
		}
		slog.Error(fmt.Sprintf("db error log. %v", get_err))
		return nil, RdbErrUnexpected
	}
	return &customer, nil
}
