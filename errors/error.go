package errors

import (
	"errors"
	"fmt"
)

// OpenFile occurs while it's unable to open specific file
func OpenFile(path string) error {
	return fmt.Errorf("unable to open file: '%v'", path)
}

// GettingValueAtKey occurs while trying to get value at a specific key in the store
func GettingValueAtKey(key string) error {
	return fmt.Errorf("error getting value at key: '%v'", key)
}

// SavingValueAtKey occurs while trying to save value at a specific key in the store
func SavingValueAtKey(key string) error {
	return fmt.Errorf("error saving value at key: '%v'", key)
}

var (
	// UnmarshalPayload occurs while trying to unmarshal request body to specific struct
	UnmarshalPayload = errors.New("error unmarshalling payload")

	// MarshalPayload occurs while trying to marshal response
	MarshalPayload = errors.New("error marshalling payload")
)
