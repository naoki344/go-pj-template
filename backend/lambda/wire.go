//+build wireinject

package main

import "github.com/google/wire"

func InitializeNoteService(cfg *Config) *noteService {
    wire.Build(SuperSet, NewNoteService)
    return nil
}
