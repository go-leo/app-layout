package app_layout

import (
	"github.com/go-leo/app-layout/adapter"
	"github.com/go-leo/app-layout/assembler"
	"github.com/go-leo/app-layout/command"
	"github.com/go-leo/app-layout/model"
	"github.com/go-leo/app-layout/query"
	"github.com/go-leo/app-layout/service"
	"github.com/go-leo/app-layout/transport"
	"github.com/google/wire"
)

var Provider = wire.NewSet(
	transport.Provider,
	assembler.Provider,
	command.Provider,
	query.Provider,
	service.Provider,
	adapter.Provider,
	model.Provider,
)
