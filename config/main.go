package config

import (
	"context"

	"github.com/gsols/go-logger"
)

var config = new(Config)

type Config struct {
	Telemetry Telemetry `mapstructure:"telemetry"`

	Vanity Vanity `mapstructure:",squash"`
}

func Get() Config {
	return *config
}

func Bootstrap(ctx context.Context, opts ...Option) {
	ctx = logger.Ctx(ctx).With().Str("process", "config").Logger().WithContext(ctx)

	for _, o := range opts {
		o(config)
	}
	Read(ctx)

	logger.Ctx(ctx).Trace().Msgf("Config loaded: %+v", config)
}
