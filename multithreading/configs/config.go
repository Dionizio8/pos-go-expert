package configs

import (
	"github.com/spf13/viper"
)

var cfg *conf

type conf struct {
	APIURLBrasilApi string `mapstructure:"API_URL_BRASIL_API"`
	APIURLViacep    string `mapstructure:"API_URL_VIA_CEP"`
}

func LoadConfig(path string) (*conf, error) {
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

	return cfg, err
}
