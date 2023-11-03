package rdbadapter

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
)


type Customer struct {
	ID      int64  `bun:"id"`
	Title   string `bun:"title"`
	Content string `bun:"content"`
}


func (rdb *MySQL) GetCustomerByID(customerID int64) (Customer, error){
	db := bun.NewDB(rdb.Conn, mysqldialect.New())
	customer := Customer{}
	get_err := db.NewSelect().Model(&customer).Where("? = ?", bun.Ident("id"), customerID).Scan(context.Background())
	if get_err != nil {
		return Customer{}, nil
	}
	return customer, nil
}
