package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
)

type MySQLConfig struct {
	HOST     string `mapstructure:"host"`
	PORT     string `mapstructure:"port"`
	USER     string `mapstructure:"username"`
	PASSWORD string `mapstructure:"password"`
	DBNAME   string `mapstructure:"dbname"`
}
type Config struct {
	MYSQL MySQLConfig `mapstructure:"mysql"`
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
