package config

type Server struct {
	HSOT string `mapstructure:"host"`
	PORT string `mapstructure:"port"`
}
