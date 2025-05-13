package config

type RemoteConfig struct {
	REMOTEDIR string `mapstructure:"remote_dir"`
	ADBPATH   string `mapstructure:"adb_path"`
}
