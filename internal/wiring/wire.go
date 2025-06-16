//go:build wireinject
// +build wireinject

//
//go:generate go run github.com/google/wire/cmd/wire
package wiring

import (
	"github.com/google/wire"
	"github.com/upinmcSE/GoLoad/internal/configs"
	"github.com/upinmcSE/GoLoad/internal/dataaccess"
	"github.com/upinmcSE/GoLoad/internal/handler"
	"github.com/upinmcSE/GoLoad/internal/handler/grpc"
	"github.com/upinmcSE/GoLoad/internal/logic"
)

var WireSet = wire.NewSet(
	configs.WireSet,
	dataaccess.WireSet,
	logic.WireSet,
	handler.WireSet,
)

func InitializeStandaloneServer(configFilePath configs.ConfigFilePath) (grpc.Server, func(), error) {
	wire.Build(WireSet)

	return nil, nil, nil
}