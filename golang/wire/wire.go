//go:build wireinject
// +build wireinject

package main

import (
	"golang/wire/internal/service"

	"github.com/google/wire"
)

func InitEvent(pp string) (service.Event, func(), error) {
	panic(wire.Build(service.NewGreeter, service.NewEvent, service.NewMessage))
}
