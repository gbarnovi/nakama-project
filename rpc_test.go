package main

import (
	"context"
	"encoding/json"
	"github.com/heroiclabs/nakama-common/runtime"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	rpcError "nakama-project/errors"
	"nakama-project/mocks"
	"reflect"
	"testing"
)

type mockStorage struct{}

func (m *mockStorage) checkIfKeyExists(ctx context.Context, nk runtime.NakamaModule, key string) (bool, error) {
	return false, nil
}

func (m *mockStorage) save(ctx context.Context, nk runtime.NakamaModule, data StorageEntry) error {
	return nil
}

type testCase struct {
	name        string
	mocksToCall func()
	input       interface{}
	expected    interface{}
	wantErr     error
}

func Test_configRPC_GetConfig(t *testing.T) {
	module := mocks.NewNakamaModule(t)
	logger := mocks.NewRuntimeLogger(t)
	c := &configRPC{
		storage: &mockStorage{},
	}

	tests := []testCase{
		{
			name: "unable to open file",
			input: Request{
				Type: "core1",
			},
			mocksToCall: func() {
				logger.On("Error", mock.Anything).Return(nil)
			},
			expected: "",
			wantErr:  rpcError.OpenFile("core1/1.0.0"),
		},
		{
			name:  "request with no payload",
			input: "",
			mocksToCall: func() {
				logger.On("Error", mock.Anything).Return(nil)
			},
			expected: "",
			wantErr:  rpcError.UnmarshalPayload,
		},
		{
			name:        "request with empty payload",
			input:       Request{},
			mocksToCall: func() {},
			expected: Response{
				Type:    "core",
				Version: "1.0.0",
				Hash:    "224d9a1f4a8504e296f4f7883a3806fc26c0eb010b0a89eb750f58ef14aa16e4",
				Content: "{\n  \"value\": \"from-core\"\n}",
			},
			wantErr: nil,
		},
		{
			name: "request with mismatched hash",
			input: Request{
				Type:    "core",
				Version: "1.0.0",
				Hash:    "some dummy hash",
			},
			mocksToCall: func() {},
			expected: Response{
				Type:    "core",
				Version: "1.0.0",
				Hash:    "224d9a1f4a8504e296f4f7883a3806fc26c0eb010b0a89eb750f58ef14aa16e4",
				Content: "null",
			},
			wantErr: nil,
		},
		{
			name: "request with matching hashes",
			input: Request{
				Type:    "core",
				Version: "1.0.0",
				Hash:    "224d9a1f4a8504e296f4f7883a3806fc26c0eb010b0a89eb750f58ef14aa16e4",
			},
			mocksToCall: func() {},
			expected: Response{
				Type:    "core",
				Version: "1.0.0",
				Hash:    "224d9a1f4a8504e296f4f7883a3806fc26c0eb010b0a89eb750f58ef14aa16e4",
				Content: "{\n  \"value\": \"from-core\"\n}",
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mocksToCall()
			encoded, err := json.Marshal(tt.input)
			if err != nil {
				t.Errorf("an error occurred while marshalling test input, %v", err.Error())
			}

			got, err := c.GetConfig(context.Background(), logger, nil, module, string(encoded))

			if tt.wantErr != nil {
				if !assert.EqualError(t, tt.wantErr, err.Error()) {
					t.Errorf("handler returned wrong message for error: got %v expected %v",
						err.Error(), tt.wantErr)
				}
			} else {
				response := Response{}
				err := json.Unmarshal([]byte(got), &response)
				if err != nil {
					t.Errorf("an error occurred while unmarshalling respose, %v", err.Error())
				}

				if !reflect.DeepEqual(tt.expected, response) {
					t.Errorf("\n got: %v, \n wanted: %v", response, tt.expected)
				}
			}
		})
	}
}
