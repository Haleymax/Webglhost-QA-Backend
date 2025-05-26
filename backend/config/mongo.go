package config

type MongoConfig struct {
	HOST string `mapstructure:"host"`
	PORT int    `mapstructure:"port"`
	USER string `mapstructure:"user"`
	PWD  string `mapstructure:"pwd"`
}
