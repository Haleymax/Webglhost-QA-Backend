package config

type FeishuConfig struct {
	FORMAPI      string `mapstructure:"form_api"`
	USERTOKENAPI string `mapstructure:"user_token_api"`
	SHEETTOKEN   string `mapstructure:"sheetToken"`
	TAPSHEETID   string `mapstructure:"tap_sheetId"`
	APPID        string `mapstructure:"app_id"`
	APPSECRET    string `mapstructure:"app_secret"`
	RANGE        string `mapstructure:"range"`
}
