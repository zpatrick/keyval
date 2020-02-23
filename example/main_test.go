package main

import (
	"testing"

	"github.com/zpatrick/keyval"
)

func TestThing(t *testing.T) {
	schema := keyval.Schema{
		IntSchemas: []*keyval.IntSchema{
			{
				Key:       "redis.port",
				Default:   4000,
				Validate:  keyval.ValidateIntBetween(0, 65535),
				Providers: []keyval.IntProvider{},
			},
		},
		StringSchemas: map[string]keyval.StringSchema{},
	}
	// foo := NewTestFoo(
	// 	keyval.WithInt("redis.port", 8080),
	// 	keyval.WithString("redis.host", "localhost"),
	// )

	// if err := thing(foo); err != nil {
	// 	t.Fatal(err)
	// }

	t.Log(schema)
}
