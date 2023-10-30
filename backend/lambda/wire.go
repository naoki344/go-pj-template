//+build wireinject

package main

import (
	"github.com/google/wire"
	"log/slog"
)

func InitializeEnAPIService(cfg *Config, logger *slog.Logger) *enAPIService {
    wire.Build(SuperSet, NewEnAPIService)
    return nil
}
