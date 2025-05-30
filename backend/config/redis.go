package config

type RedisConfig struct {
	HOST string `mapstructure:"host"`
	PORT int    `mapstructure:"port"`
	DB   int    `mapstructure:"db"`
}
