package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		Bot `yaml:"bot"`
	}

	Bot struct {
		Name     string   `env-required:"true" yaml:"name"  env:"BLOSSOM_TW_BOT_NAME"`
		OAuth    string   `env-required:"true" yaml:"oauth" env:"BLOSSOM_TW_BOT_OAUTH"`
		LogLevel string   `env-required:"true" yaml:"log_level" env:"BLOSSOM_TW_LOG_LEVEL"`
		Channel  []string `env-required:"true" yaml:"channel"  env:"BLOSSOM_TW_CHANNEL" env-delim:","`
	}
)

// New returns app config.
func New(configPath string) (*Config, error) {
	c := &Config{}

	err := cleanenv.ReadConfig(configPath, c)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
