package config

import (
	"github.com/ducka/demoroo/internal/grpc"
	"go.uber.org/fx"
)

// LoadConfig Loads a configuration object populated from environmental variables
func LoadConfig[TConfig interface{}](prefix string, config TConfig) (TConfig, error) {

	//
	// NB: This is just mock code. Here you'd be loading your config from environmental variables
	switch any(config).(type) {
	case grpc.ServerSettings:
		c := any(&config).(*grpc.ServerSettings)
		c.HostAddress = "localhost:8080"
	}
	//
	//

	return config, nil
}

// RegisterConfig Returns a provider function that loads a configuration object. This is to be used with
// dependency injection libraries like uber fx.
func RegisterConfig[TConfig interface{}](prefix string, config TConfig) fx.Option {
	return fx.Provide(
		func() (TConfig, error) {
			return LoadConfig[TConfig](prefix, config)
		},
	)
}
