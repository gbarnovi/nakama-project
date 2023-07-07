package main

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/heroiclabs/nakama-common/runtime"
	rpcError "nakama-project/errors"
	"os"
	"strings"
)

const (
	DefaultType    string = "core"
	DefaultVersion string = "1.0.0"
	DefaultHash    string = "null"
)

type configRPC struct {
	storage Store
}

// GetConfig returns configuration from file
func (c *configRPC) GetConfig(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {
	requestBody := &Request{
		Type:    DefaultType,
		Version: DefaultVersion,
		Hash:    DefaultHash,
	}

	if err := json.Unmarshal([]byte(payload), requestBody); err != nil {
		logger.Error(fmt.Sprintf("error unmarshalling payload, err: %v", err))
		return "", rpcError.UnmarshalPayload
	}

	key := fmt.Sprintf("%v/%v", strings.ToLower(requestBody.Type), strings.ToLower(requestBody.Version))

	content, err := os.ReadFile(fmt.Sprintf("./data/%s/%s.json", requestBody.Type, requestBody.Version))
	if err != nil {
		logger.Error(fmt.Sprintf("error opening file: '%s/%s', err: %v", requestBody.Type, requestBody.Version, err))
		return "", rpcError.OpenFile(key)
	}

	hash := sha256.Sum256(content)

	if requestBody.Hash != "null" && requestBody.Hash != fmt.Sprintf("%x", hash) {
		content = []byte("null")
	}

	ok, err := c.storage.checkIfKeyExists(ctx, nk, key)
	if err != nil {
		logger.Error(fmt.Sprintf("error getting value at key '%v', err: %v", key, err))
		return "", rpcError.GettingValueAtKey(key)
	}

	storageEntry := StorageEntry{
		Type:    requestBody.Type,
		Version: requestBody.Version,
		Hash:    fmt.Sprintf("%x", hash),
		Content: string(content),
	}

	if !ok {
		err = c.storage.save(ctx, nk, storageEntry)
		if err != nil {
			logger.Error(fmt.Sprintf("error saving value at key: '%v', err: %v", key, err))
			return "", rpcError.SavingValueAtKey(key)
		}
	}

	response := Response{
		Type:    requestBody.Type,
		Version: requestBody.Version,
		Hash:    fmt.Sprintf("%x", hash),
		Content: string(content),
	}

	bytes, err := json.Marshal(response)
	if err != nil {
		logger.Error(fmt.Sprintf("error marshalling payload, err %v", err))
		return "", rpcError.MarshalPayload
	}

	return string(bytes), nil
}
