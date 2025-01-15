package config

import (
	"github.com/spf13/viper"
	"log"
	"strings"
)

type Config struct {
	Couchbase     Couchbase
	Elasticsearch Elasticsearch
	Server        ServerConfig
}

type ServerConfig struct {
	Host string
	Port string
}

type Couchbase struct {
	Host     string
	UserName string
	Password string
	Bucket   string
}

type Elasticsearch struct {
	Url string
}

func LoadConfig(env string) *Config {
	viper.SetConfigName("config-" + env)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("internal/config")

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Config could not be load: %v", err)
	}
	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("Config could not be parsed: %v", err)
	}

	return &cfg
}
