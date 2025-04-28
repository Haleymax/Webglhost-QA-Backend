package config

type MySQLConfig struct {
	HOST     string `mapstructure:"host"`
	PORT     string `mapstructure:"port"`
	USER     string `mapstructure:"username"`
	PASSWORD string `mapstructure:"password"`
	DBNAME   string `mapstructure:"dbname"`
}
