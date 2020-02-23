package keyval

import (
	"fmt"
)

type KeyNotDefinedError struct {
	Key string
}

func NewKeyNotDefinedError(key string) *KeyNotDefinedError {
	return &KeyNotDefinedError{key}
}

func (e *KeyNotDefinedError) Error() string {
	return fmt.Sprintf("%s: key not defined", e.Key)
}

type KeyNotFoundError struct {
	Key string
}

func NewKeyNotFoundError(key string) *KeyNotFoundError {
	return &KeyNotFoundError{key}
}

func (e *KeyNotFoundError) Error() string {
	return fmt.Sprintf("%s: key not found", e.Key)
}
