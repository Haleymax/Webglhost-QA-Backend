package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	MYSQL  MySQLConfig  `mapstructure:"mysql"`
	SERVER Server       `mapstructure:"server"`
	REMOTE RemoteConfig `mapstructure:"remote"`
	MONGO  MongoConfig  `mapstructure:"mongo"`
	REDIS  RedisConfig  `mapstructure:"redis"`
	FEISHU FeishuConfig `mapstructure:"feishu"`
}

func LoadConfig() *Config {
	workDir, _ := os.Getwd()
	configFilePath := filepath.Join(workDir, "config", "config.yaml")
	log.Println("Loading config from", configFilePath)

	viper.SetConfigFile(configFilePath)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	log.Println("Using config file:", viper.AllSettings())

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("unable to decode config, %v", err)
	}
	return &cfg
}
