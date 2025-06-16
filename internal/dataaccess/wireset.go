package dataaccess

import (
	"github.com/google/wire"
	"github.com/upinmcSE/GoLoad/internal/dataaccess/database"
)

var WireSet = wire.NewSet(
	database.WireSet,
)