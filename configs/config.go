package configs

import (
	"os"
	"time"

	"github.com/spf13/viper"
	"goparking/internals/libs/logger"
)

const (
	ProductionEnv   = "production"
	DatabaseTimeout = time.Second * 5
)

type Config struct {
	Environment    string `mapstructure:"ENVIRONMENT"`
	HttpPort       int    `mapstructure:"HTTP_PORT"`
	AuthSecret     string `mapstructure:"AUTH_SECRET"`
	DatabaseURI    string `mapstructure:"DATABASE_URI"`
	MinioEndpoint  string `mapstructure:"MINIO_ENDPOINT"`
	MinioAccessKey string `mapstructure:"MINIO_ACCESSKEY"`
	MinioSecretKey string `mapstructure:"MINIO_SECRETKEY"`
	MinioBucket    string `mapstructure:"MINIO_BUCKET"`
	MinioBaseurl   string `mapstructure:"MINIO_BASEURL"`
	MinioUseSSL    bool   `mapstructure:"MINIO_USESSL"`
	RedisURI       string `mapstructure:"REDIS_URI"`
	RedisPassword  string `mapstructure:"REDIS_PASSWORD"`
	RedisDB        int    `mapstructure:"REDIS_DB"`
}

var (
	cfg Config
)

func LoadConfig() *Config {
	viper.AutomaticEnv()

	if _, err := os.Stat("app.env"); err == nil {
		viper.SetConfigFile("app.env")
		viper.SetConfigType("env")
		if err := viper.ReadInConfig(); err != nil {
			logger.Fatal("Error loading configuration file: %v", err)
		}
	}

	cfg = Config{
		Environment:    viper.GetString("ENVIRONMENT"),
		HttpPort:       viper.GetInt("HTTP_PORT"),
		AuthSecret:     viper.GetString("AUTH_SECRET"),
		DatabaseURI:    viper.GetString("DATABASE_URI"),
		MinioEndpoint:  viper.GetString("MINIO_ENDPOINT"),
		MinioAccessKey: viper.GetString("MINIO_ACCESSKEY"),
		MinioSecretKey: viper.GetString("MINIO_SECRETKEY"),
		MinioBucket:    viper.GetString("MINIO_BUCKET"),
		MinioBaseurl:   viper.GetString("MINIO_BASEURL"),
		MinioUseSSL:    viper.GetBool("MINIO_USESSL"),
		RedisURI:       viper.GetString("REDIS_URI"), // ✅ Đảm bảo đọc REDIS_URI
		RedisPassword:  viper.GetString("REDIS_PASSWORD"),
		RedisDB:        viper.GetInt("REDIS_DB"),
	}

	if cfg.DatabaseURI == "" {
		logger.Fatal("DATABASE_URI is not set!")
	}

	return &cfg
}

func GetConfig() *Config {
	return &cfg
}
