package config

import (
	"fmt"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		Environment string
		Http        HttpConfig
		Limiter     LimiterConfig
		Most        MostConfig
	}

	HttpConfig struct {
		Host               string        `yaml:"http.host" env:"HOST" env-default:"localhost"`
		Port               string        `yaml:"http.port" env:"PORT" env-default:"5432"`
		ReadTimeout        time.Duration `yaml:"http.readTimeout" env:"READ_TIMEOUT" env-default:"10s"`
		WriteTimeout       time.Duration `yaml:"http.writeTimeout" env:"WRITE_TIMEOUT" env-default:"10s"`
		MaxHeaderMegabytes int           `yaml:"http.maxHeaderBytes" env-default:"1"`
	}

	LimiterConfig struct {
		RPS   int           `yaml:"limiter.rps" env-default:"10"`
		Burst int           `yaml:"limiter.burst" env-default:"20"`
		TTL   time.Duration `yaml:"limiter.ttl" env-default:"10m"`
	}

	MostConfig struct {
		Server    string `env:"MOST_SERVER"`
		Token     string `env:"MOST_TOKEN"`
		ChannelId string `env:"MOST_CHANNEL_ID"`
	}
)

func Init(configDir string) (*Config, error) {
	var conf Config

	if err := cleanenv.ReadConfig(configDir, &conf); err != nil {
		return nil, fmt.Errorf("failed to read config file. error: %w", err)
	}

	if err := cleanenv.ReadEnv(&conf); err != nil {
		return nil, fmt.Errorf("failed to read env file. error: %w", err)
	}

	return &conf, nil
}
