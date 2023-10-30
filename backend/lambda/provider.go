package main

import (
    "github.com/google/wire"
)

var SuperSet = wire.NewSet(
    NewDB,
    NewGetCustomerByIDRepository,
)
