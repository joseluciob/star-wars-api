package configs

import (
	"log"

	"github.com/spf13/viper"
)

type Configs struct {
	AppPrefix string `mapstructure:"app_prefix"`
	Port      string `mapstructure:"http_port"`
}

func New() (*Configs, error) {
	viper.SetDefault("HTTP_PORT", "80")
	viper.SetConfigType("env")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Println("config: .env file not found")
	}

	cfg := &Configs{}
	if err := viper.Unmarshal(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
