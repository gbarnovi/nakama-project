package main

import "github.com/heroiclabs/nakama-common/runtime"

//go:generate mockery --name NakamaModule --case=underscore
type NakamaModule interface {
	runtime.NakamaModule
}

//go:generate mockery --name RuntimeLogger --case=underscore
type RuntimeLogger interface {
	runtime.Logger
}
