package config

import (
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		Bot `yaml:"bot"`
		AI  `yaml:"ai"`
	}

	Bot struct {
		Name            string        `env-required:"true" yaml:"name"  env:"BLOSSOM_TW_BOT_NAME"`
		OAuth           string        `env-required:"true" yaml:"oauth" env:"BLOSSOM_TW_BOT_OAUTH"`
		LogLevel        string        `env-required:"true" yaml:"log_level" env:"BLOSSOM_TW_LOG_LEVEL"`
		Channel         []string      `env-required:"true" yaml:"channel"  env:"BLOSSOM_TW_CHANNEL" env-delim:","`
		IgnoreChannels  []string      `env-required:"true" yaml:"ignore_channels"  env:"BLOSSOM_TW_IGNORE_CHANNELS" env-delim:","`
		CommandsEnabled []string      `env-required:"true" yaml:"commands_enabled"  env:"BLOSSOM_TW_COMMANDS_ENABLED" env-delim:","`
		CmdTimeout      time.Duration `env-required:"false" env-default:"20s" yaml:"cmd_timeout" env:"BLOSSOM_TW_CMD_TIMEOUT"`
	}

	AI struct {
		OpenAiApiKey       string `env-required:"true" yaml:"openai_api_key" env:"BLOSSOM_TW_AI_API_KEY"`
		MaxTokens          int    `env-required:"true" yaml:"max_tokens" env:"BLOSSOM_TW_AI_MAX_TOKENS"`
		CustomInstructions string `env-required:"false" yaml:"custom_instructions" env:"BLOSSOM_TW_AI_CUSTOM_INSTRUCTIONS"`
		MaxReplyLen        int    `env-required:"false" env-default:"128" yaml:"max_reply_len" env:"BLOSSOM_TW_AI_MAX_REPLY_LEN"`
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
