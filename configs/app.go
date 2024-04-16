package configs

import "github.com/spf13/viper"

var cfg *AppConfig

type AppConfig struct {
	WebServerBaseUrl string `mapstructure:"WEB_SERVER_BASE_URL"`
	WebServerPort    string `mapstructure:"WEB_SERVER_PORT"`
}

func LoadConfig(path string) (*AppConfig, error) {
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
	return cfg, nil
}
