/*
Package yourpackage does something interesting.
*/
package main

import (
	"github.com/g-stayfresh/en/backend/api/lib/ogen"
	"log/slog"
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


type enAPIService struct {
	dbRepository *GetCustomerByIDRepository
	logger *slog.Logger
}

func (n *enAPIService) PostCreateCustomer(ctx context.Context, req *ogen.PostCreateCustomerReq) (*ogen.PostCreateCustomerOK, error) {
	return &ogen.PostCreateCustomerOK{}, nil
}

func (n *enAPIService) GetCustomerByID(ctx context.Context, params ogen.GetCustomerByIDParams) (*ogen.GetCustomerByIDOK, error) {
	n.logger.Error("service unavailable.")
	db := bun.NewDB(n.dbRepository.db.conn, mysqldialect.New())

	customer_db := Customer{}
	get_err := db.NewSelect().Model(&customer_db).Where("? = ?", bun.Ident("id"), params.CustomerID).Scan(context.Background())
	if get_err != nil {
		panic(get_err)
	}
	customer := ogen.GetCustomerByIDOK{
		ID:      customer_db.ID,
		Title:   customer_db.Title,
		Content: customer_db.Content,
	}

	return &customer, nil
}


func NewEnAPIService(dbRepository *GetCustomerByIDRepository, logger *slog.Logger) *enAPIService {
	return &enAPIService{
		dbRepository: dbRepository,
		logger: logger,
	}
}
