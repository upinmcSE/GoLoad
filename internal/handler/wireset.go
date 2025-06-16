package handler

import (
	"github.com/google/wire"
	"github.com/upinmcSE/GoLoad/internal/handler/grpc"
	"github.com/upinmcSE/GoLoad/internal/handler/http"
)

var WireSet = wire.NewSet(
	grpc.WireSet,
	http.WireSet,
)