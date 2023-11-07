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
	getCustomerByIDPort := rdbport.NewGetCustomerByIDPort(db)
	getCustomerByIDUsecase := customerusecase.NewGetCustomerByIDUsecase(getCustomerByIDPort)
	getCustomerByIDAPIPort := apiport.NewGetCustomerByIDAPIPort(getCustomerByIDUsecase)
	enAPIAdapter := ogenadapter.NewEnAPIAdapter(getCustomerByIDAPIPort)
	return enAPIAdapter
}