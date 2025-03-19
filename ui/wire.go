package ui

import (
	"github.com/go-leo/app-layout/api/v1"
	"github.com/google/wire"
)

var Provider = wire.NewSet(
	NewUserService, wire.Bind(new(api.UserService), new(*UserService)),
)
