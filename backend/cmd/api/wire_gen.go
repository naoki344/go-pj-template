// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/g-stayfresh/en/backend/internal/adapter/driven/rdb"
	"github.com/g-stayfresh/en/backend/internal/adapter/driver/ogen"
	"github.com/g-stayfresh/en/backend/internal/port/driven/rdb"
	"github.com/g-stayfresh/en/backend/internal/port/driver/api"
	"github.com/g-stayfresh/en/backend/internal/usecase/customer"
)

import (
	_ "github.com/go-sql-driver/mysql"
)

// Injectors from wire.go:

func InitializeEnAPIService(db *rdbadapter.MySQL) *ogenadapter.EnAPIAdapter {
	customerRdbPort := rdbport.NewCustomerRdbPort(db)
	customerUsecase := customerusecase.NewCustomerUsecase(customerRdbPort)
	customerAPIPort := apiport.NewCustomerAPIPort(customerUsecase)
	enAPIAdapter := ogenadapter.NewEnAPIAdapter(customerAPIPort)
	return enAPIAdapter
}
