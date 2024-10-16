package config

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Env       string        `mapstructure:"env"`
	Port      string        `mapstructure:"port"`
	Timeout   time.Duration `mapstructure:"timeout"`
	Product   ProductGRPC   `mapstructure:"grpcProduct"`
	Auth      AuthGRPC      `mapstructure:"grpcAuth"`
	SecretKey string        `mapstructure:"secretKey"`
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

// read config from ./config/
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

// set SECRET_KEY env variable
func SetEnvSecret(secret string) error {

	err := os.Setenv("SECRET_KEY", secret)

	return err
}
