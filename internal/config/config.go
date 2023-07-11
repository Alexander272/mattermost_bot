package config

import (
	"fmt"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		Environment string `env-default:"dev"`
		Http        HttpConfig
		Limiter     LimiterConfig
		Most        MostConfig
	}

	HttpConfig struct {
		Host               string        `yaml:"host" env:"HOST" env-default:"localhost"`
		Port               string        `yaml:"port" env:"PORT" env-default:"11000"`
		ReadTimeout        time.Duration `yaml:"readTimeout" env:"READ_TIMEOUT" env-default:"10s"`
		WriteTimeout       time.Duration `yaml:"writeTimeout" env:"WRITE_TIMEOUT" env-default:"10s"`
		MaxHeaderMegabytes int           `yaml:"maxHeaderBytes" env-default:"1"`
	}

	LimiterConfig struct {
		RPS   int           `yaml:"rps" env-default:"10"`
		Burst int           `yaml:"burst" env-default:"20"`
		TTL   time.Duration `yaml:"ttl" env-default:"10m"`
	}

	MostConfig struct {
		Server    string `env:"MOST_SERVER"`
		Token     string `env:"MOST_TOKEN"`
		ChannelId string `env:"MOST_CHANNEL_ID"`
	}
)

func Init(path string) (*Config, error) {
	var conf Config

	if err := cleanenv.ReadConfig(path, &conf); err != nil {
		return nil, fmt.Errorf("failed to read config file. error: %w", err)
	}

	if err := cleanenv.ReadEnv(&conf); err != nil {
		return nil, fmt.Errorf("failed to read env file. error: %w", err)
	}

	return &conf, nil
}
