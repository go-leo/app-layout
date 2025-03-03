package app_layout

import (
	"github.com/go-leo/app-layout/adapter"
	"github.com/go-leo/app-layout/application"
	"github.com/go-leo/app-layout/assembler"
	"github.com/go-leo/app-layout/model"
	"github.com/go-leo/app-layout/transport"
	"github.com/google/wire"
)

var Provider = wire.NewSet(
	transport.Provider,
	application.Provider,
	assembler.Provider,
	adapter.Provider,
	model.Provider,
)
