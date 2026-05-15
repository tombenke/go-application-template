package application

import (
	"github.com/spf13/pflag"
	"github.com/tombenke/go-12f-common/v2/config"
	"github.com/tombenke/go-application-template/internal/infrastructure/webserver"
	"go.uber.org/multierr"
)

// The configuration parameters of the application.
type Config struct {
	webserver webserver.Config
}

// Add application-specific config parameters to flagset.
func (cfg *Config) GetConfigFlagSet(fs *pflag.FlagSet) {
	cfg.webserver.GetConfigFlagSet(fs)
}

func (cfg *Config) LoadConfig(fs *pflag.FlagSet) error {
	return multierr.Combine(
		cfg.webserver.LoadConfig(fs),
	)
}

var _ config.Configurer = (*Config)(nil)
