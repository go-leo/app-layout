package application

import (
	"github.com/go-leo/app-layout/application/command"
	"github.com/go-leo/app-layout/application/query"
	"github.com/go-leo/app-layout/application/service"
	"github.com/google/wire"
)

var Provider = wire.NewSet(
	command.Provider,
	query.Provider,
	service.Provider,
)
