package main

import "github.com/zpatrick/keyval"

func main() {
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

	// env := keyval.NewEnvironmentProvider()
	// file := keyval.NewFileProvider("config.yaml", keyval.YAMLParser)
	// flags := keyval.NewFlagsProvider("", flag.ContinueOnError)

	// foo := keyval.Foo{
	// 	IntValues: map[string]keyval.IntValue{
	// 		"redis.port": {
	// 			Default:  4000,
	// 			Validate: keyval.ValidateIntBetween(0, 65535),
	// 			Providers: []keyval.IntProvider{
	// 				flags.Int("redis-port", 4000, "redis port"),
	// 				env.Int("APP_REDIS_PORT"),
	// 				file.Int("redis", "port"),
	// 			},
	// 		},
	// 	},
	// }

	// if err := foo.Validate(); err != nil {
	// 	log.Fatal(err)
	// }

	// if err := flags.Parse(os.Args); err != nil {
	// 	log.Fatal(err)
	// }

	// redisPort, err := foo.Int("redis.port")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// exists, err := foo.Exists("redis.host")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// redisClient := redis.Client{
	// 	Host:    foo.MustString("redis.host"),
	// 	Port:    foo.MustInt("redis.port"),
	// 	Timeout: foo.MustDuration("redis.timeout"),
	// }

}
