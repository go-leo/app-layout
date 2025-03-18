package main

import (
	"github.com/go-leo/app-layout/acl"
	"github.com/go-leo/app-layout/handler"
	"github.com/go-leo/app-layout/infra"
	"github.com/go-leo/app-layout/model"
	"github.com/go-leo/app-layout/ui"
	"github.com/google/wire"
)

var Provider = wire.NewSet(
	cmd.Provider,
	cq.Provider,
	model.Provider,
	acl.Provider,
	infra.Provider,
)
