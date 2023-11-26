//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	rdbadapter "github.com/naoki344/go-pj-template/backend/internal/adapter/driven/rdb"
	ogenadapter "github.com/naoki344/go-pj-template/backend/internal/adapter/driver/ogen"
	rdbport "github.com/naoki344/go-pj-template/backend/internal/port/driven/rdb"
	apiport "github.com/naoki344/go-pj-template/backend/internal/port/driver/api"
	customerusecase "github.com/naoki344/go-pj-template/backend/internal/usecase/customer"
)

func InitializeAPIService(db *rdbadapter.MySQL) *ogenadapter.APIAdapter {
	wire.Build(
		ogenadapter.NewAPIAdapter,
		apiport.NewCustomerAPIPort,
		customerusecase.NewCustomerUsecase,
		wire.Bind(new(apiport.CustomerAPIPortInterface), new(*apiport.CustomerAPIPort)),
		wire.Bind(new(customerusecase.CustomerUsecaseInterface), new(*customerusecase.CustomerUsecase)),
		rdbport.NewRdbPort,
		wire.Bind(new(rdbport.RdbPortInterface), new(*rdbport.RdbPort)),
		// NewMySQLはmain内に実装するため、bindのみ行う
		wire.Bind(new(rdbadapter.RdbInterface), new(*rdbadapter.MySQL)),
	)
	return nil
}
