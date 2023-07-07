package main

import (
	"context"
	"database/sql"
	"github.com/heroiclabs/nakama-common/runtime"
)

func InitModule(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	configRPC := configRPC{storage: &storage{}}

	if err := initializer.RegisterRpc("custom_rpc_function", configRPC.GetConfig); err != nil {
		logger.Error(`error registering "custom_rpc_function" rpc`)
		return err
	}

	return nil
}
