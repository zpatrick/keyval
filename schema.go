package keyval

import (
	"errors"
	"fmt"
)

type IntSchema struct {
	Key           string
	Default       int
	DefaultIsZero bool
	Validate      IntValidator
	Providers     []IntProvider
}

func (s *IntSchema) Value() (int, error) {
	var value *int
	for _, p := range s.Providers {
		v, err := p.Int(s.Key)
		if err != nil {
			if knf := new(KeyNotFoundError); errors.Is(err, knf) {
				continue
			}

			return 0, fmt.Errorf("%s: failed to lookup: %w", s.Key, err)
		}

		value = &v
		break
	}

	if value == nil {
		if !s.DefaultIsZero && s.Default == 0 {
			return 0, NewKeyNotFoundError(s.Key)
		}

		value = &s.Default
	}

	if err := s.Validate(*value); err != nil {
		return 0, fmt.Errorf("%s: failed validation: %w", s.Key, err)
	}

	return *value, nil
}

type StringSchema struct {
	Key           string
	Default       string
	DefaultIsZero bool
	Validate      StringValidator
	Providers     []StringProvider
}

func (s *StringSchema) Value() (string, error) {
	var value *string
	for _, p := range s.Providers {
		v, err := p.String(s.Key)
		if err != nil {
			if knf := new(KeyNotFoundError); errors.Is(err, knf) {
				continue
			}

			return "", fmt.Errorf("%s: failed to lookup int: %w", s.Key, err)
		}

		value = &v
		break
	}

	if value == nil {
		if !s.DefaultIsZero && s.Default == "" {
			return "", NewKeyNotFoundError(s.Key)
		}

		value = &s.Default
	}

	if err := s.Validate(*value); err != nil {
		return "", fmt.Errorf("%s: failed validation: %w", s.Key, err)
	}

	return *value, nil
}

type Schema struct {
	IntSchemas    []*IntSchema
	StringSchemas map[string]StringSchema
}

func (s *Schema) Validate() error {
	// no duplicate keys
	// no key has 0 providers
	// no key is nil or has nil providers
	// no key is empty string
	// should we get real values and call validate on each? add param to this method?
}
