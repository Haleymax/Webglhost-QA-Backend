package config

type MySQLConfig struct {
	HOST     string `mapstructure:"host"`
	PORT     string `mapstructure:"port"`
	USER     string `mapstructure:"user"`
	PASSWORD string `mapstructure:"password"`
	DBNAME   string `mapstructure:"dbname"`
}
