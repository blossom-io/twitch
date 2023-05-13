package config

import (
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
		Ignore   []string `env-required:"true" yaml:"ignore"  env:"BLOSSOM_TW_IGNORE" env-delim:","`
	}
)

// New returns app config.
func New(configPath string) (*Config, error) {
	c := &Config{}

	err := cleanenv.ReadEnv(c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
