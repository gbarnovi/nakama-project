package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/heroiclabs/nakama-common/runtime"
	"strings"
)

type StorageEntry struct {
	Type    string `json:"type"`
	Version string `json:"version"`
	Hash    string `json:"hash"`
	Content string `json:"content"`
}

type storage struct{}

type Store interface {
	checkIfKeyExists(ctx context.Context, nk runtime.NakamaModule, key string) (bool, error)
	save(ctx context.Context, nk runtime.NakamaModule, data StorageEntry) error
}

// checkIfKeyExists checks if there is an entry for the user in the storage
func (c *storage) checkIfKeyExists(ctx context.Context, nk runtime.NakamaModule, key string) (bool, error) {
	userId := ctx.Value(runtime.RUNTIME_CTX_USER_ID).(string)

	object, err := nk.StorageRead(ctx, []*runtime.StorageRead{
		{
			Collection: "config",
			Key:        key,
			UserID:     userId,
		},
	})

	if err != nil {
		return false, err
	}

	return len(object) == 1, nil
}

// save stores a value at a specific key in the storage
func (c *storage) save(ctx context.Context, nk runtime.NakamaModule, data StorageEntry) error {
	encoded, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("an error occurred while marshalling data for storage, err: %v", err)
	}

	userId := ctx.Value(runtime.RUNTIME_CTX_USER_ID).(string)

	objects := []*runtime.StorageWrite{
		{
			Collection:      "config",
			Key:             fmt.Sprintf("%v/%v", strings.ToLower(data.Type), strings.ToLower(data.Version)),
			UserID:          userId,
			Value:           string(encoded),
			PermissionRead:  2,
			PermissionWrite: 1,
		},
	}

	if _, err := nk.StorageWrite(ctx, objects); err != nil {
		return fmt.Errorf("an error occured while saving to storage, err: %v", err)
	}

	return nil
}
