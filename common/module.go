package common

import (
	"go.uber.org/fx"

	"patient/common/cache"
	"patient/common/config"
)

var Module = fx.Module("common",
	config.Module,
	cache.Module,
	//store.Module,
)
