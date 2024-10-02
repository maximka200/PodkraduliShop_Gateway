package config

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Env         string        `mapstructure:"env"`
	Port        string        `mapstructure:"port"`
	Timeout     time.Duration `mapstructure:"timeout"`
	ProductGRPC `mapstructure:"productGRPC"`
	AuthGRPC    `mapstructure:"authGRPC"`
}

type ProductGRPC struct {
	Addr       string        `mapstructure:"addr"`
	Timeout    time.Duration `mapstructure:"timeout"`
	RetryCount int           `mapstructure:"retryCount"`
}

type AuthGRPC struct {
	Addr       string        `mapstructure:"addr"`
	Timeout    time.Duration `mapstructure:"timeout"`
	RetryCount int           `mapstructure:"retryCount"`
}

func MustReadConfig() Config {
	var cfg Config

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config/")

	if err := viper.ReadInConfig(); err != nil {
		slog.Error(err.Error())
		panic(fmt.Errorf("error reading config file: %w", err))
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		panic(fmt.Errorf("unable to unmarshal into struct: %v", err))
	}

	return cfg
}
