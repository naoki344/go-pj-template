//+build wireinject

package main

import (
	"github.com/google/wire"
	"log/slog"
)

func InitializeNoteService(cfg *Config, logger *slog.Logger) *noteService {
    wire.Build(SuperSet, NewNoteService)
    return nil
}
