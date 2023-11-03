//+build wireinject

package main

import (
	"github.com/google/wire"
	apiport "github.com/g-stayfresh/en/backend/internal/port/driver/api"
	ogenadapter "github.com/g-stayfresh/en/backend/internal/adapter/driver/ogen"
	customerusecase "github.com/g-stayfresh/en/backend/internal/usecase/customer"
	rdbport "github.com/g-stayfresh/en/backend/internal/port/driven/rdb"
	rdbadapter "github.com/g-stayfresh/en/backend/internal/adapter/driven/rdb"
)

func InitializeEnAPIService(db *rdbadapter.MySQL) *ogenadapter.EnAPIAdapter {
    wire.Build(
		ogenadapter.NewEnAPIAdapter,
		apiport.NewGetCustomerByIDAPIPort,
		customerusecase.NewGetCustomerByIDUsecase,
		wire.Bind(new(customerusecase.GetCustomerByIDInterface), new(*customerusecase.GetCustomerByIDUsecase)),
		rdbport.NewGetCustomerByIDPort,
		wire.Bind(new(rdbport.GetCustomerByIDPortInterface), new(*rdbport.GetCustomerByIDPort)),
		// NewMySQLはmain内に実装するため、bindのみ行う
		wire.Bind(new(rdbadapter.RdbInterface), new(*rdbadapter.MySQL)),
	)
    return nil
}
