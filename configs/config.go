package configs

import (
	"time"

	"github.com/spf13/viper"
	"goparking/internals/libs/logger"
)

const (
	ProductionEnv = "production"

	DatabaseTimeout = time.Second * 5
)

type Config struct {
	Environment string `mapstructure:"ENVIRONMENT"`
	HttpPort    int    `mapstructure:"HTTP_PORT"`
	AuthSecret  string `mapstructure:"AUTH_SECRET"`
	DatabaseURI string `mapstructure:"DATABASE_URI"`
}

var (
	cfg Config
)

func LoadConfig(path string) *Config {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		logger.Fatal("Error on load configuration file, error: %v", err)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		logger.Fatalf("Error on parsing configuration file, error: %v", err)
	}

	return &cfg
}

func GetConfig() *Config {
	return &cfg
}
