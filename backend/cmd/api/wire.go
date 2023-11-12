//go:build wireinject
// +build wireinject

package main

import (
	rdbadapter "github.com/g-stayfresh/en/backend/internal/adapter/driven/rdb"
	ogenadapter "github.com/g-stayfresh/en/backend/internal/adapter/driver/ogen"
	rdbport "github.com/g-stayfresh/en/backend/internal/port/driven/rdb"
	apiport "github.com/g-stayfresh/en/backend/internal/port/driver/api"
	customerusecase "github.com/g-stayfresh/en/backend/internal/usecase/customer"
	"github.com/google/wire"
)

func InitializeEnAPIService(db *rdbadapter.MySQL) *ogenadapter.EnAPIAdapter {
	wire.Build(
		ogenadapter.NewEnAPIAdapter,
		apiport.NewCustomerAPIPort,
		customerusecase.NewCustomerUsecase,
		wire.Bind(new(customerusecase.CustomerUsecaseInterface), new(*customerusecase.CustomerUsecase)),
		rdbport.NewCustomerRdbPort,
		wire.Bind(new(rdbport.CustomerRdbPortInterface), new(*rdbport.CustomerRdbPort)),
		// NewMySQLはmain内に実装するため、bindのみ行う
		wire.Bind(new(rdbadapter.RdbInterface), new(*rdbadapter.MySQL)),
	)
	return nil
}
