package webserver

import (
	"time"

	"github.com/spf13/pflag"
	"github.com/tombenke/go-12f-common/v2/config"
)

type Config struct {
	Port         int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

func (cfg *Config) GetConfigFlagSet(fs *pflag.FlagSet) {
	fs.IntVar(&cfg.Port, "web-port", 8081, "HTTP web server port")
	fs.DurationVar(&cfg.ReadTimeout, "web-read-timeout", 5*time.Second, "HTTP read timeout")
	fs.DurationVar(&cfg.WriteTimeout, "web-write-timeout", 10*time.Second, "HTTP write timeout")
	fs.DurationVar(&cfg.IdleTimeout, "web-idle-timeout", 30*time.Second, "HTTP idle timeout")
}

func (cfg *Config) LoadConfig(_ *pflag.FlagSet) error {
	if cfg.Port == 0 {
		cfg.Port = 8081
	}
	if cfg.ReadTimeout == 0 {
		cfg.ReadTimeout = 5 * time.Second
	}
	if cfg.WriteTimeout == 0 {
		cfg.WriteTimeout = 10 * time.Second
	}
	if cfg.IdleTimeout == 0 {
		cfg.IdleTimeout = 30 * time.Second
	}
	return nil
}

var _ config.Configurer = (*Config)(nil)
