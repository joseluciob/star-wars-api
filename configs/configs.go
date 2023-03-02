package configs

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type Configs struct {
	AppPrefix string `mapstructure:"app_prefix"`
	AppDebug  bool   `mapstructure:"app_debug"`
	Port      string `mapstructure:"http_port"`
	SwApiBase string `mapstructure:"swapi_base"`
	DB        DB     `mapstructure:",squash"`
	Http      Http   `mapstructure:",squash"`
}

type Http struct {
	ReadTimeout          time.Duration `mapstructure:"http_read_timeout"`
	WriteTimeout         time.Duration `mapstructure:"http_write_timeout"`
	MaxIdleConnDuration  time.Duration `mapstructure:"http_max_idle_conn_duration"`
	DialConcurrency      int           `mapstructure:"http_dial_concurrency"`
	DialDnsCacheDuration time.Duration `mapstructure:"http_dial_dns_cache_duration"`
}

type DB struct {
	User string `mapstructure:"db_user"`
	Pass string `mapstructure:"db_password"`
	Name string `mapstructure:"db_name"`
	Host string `mapstructure:"db_host"`
	Port string `mapstructure:"db_port"`
}

func New() (*Configs, error) {
	viper.SetDefault("HTTP_PORT", "80")
	viper.SetDefault("APP_DEBUG", "true")
	viper.SetDefault("SWAPI_BASE", "https://swapi.dev/api")
	viper.SetDefault("HTTP_READ_TIMEOUT", "10s")
	viper.SetDefault("HTTP_WRITE_TIMEOUT", "10s")
	viper.SetDefault("HTTP_MAX_IDLE_CONN_DURATION", "1h")
	viper.SetDefault("HTTP_DIAL_CONCURRENCY", "4096")
	viper.SetDefault("HTTP_DIAL_DNS_CACHE_DURATION", "1h")
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
