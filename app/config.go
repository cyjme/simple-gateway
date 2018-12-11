package app

import (
	"github.com/spf13/viper"
)

var (
	Config appConfig
)

type appConfig struct {
	Http struct {
		Domain string `mapstructure:"domain"`
		Port   string `mapstructure:"port"`
	} `mapstructure:"http"`
	DB struct {
		Host               string `mapstructure:"host"`
		Port               string `mapstructure:"port"`
		User               string `mapstructure:"user"`
		Password           string `mapstructure:"password"`
		Name               string `mapstructure:"name"`
		MaxIdleConnections int    `mapstructure:"max_idle_connections"`
		MaxOpenConnections int    `mapstructure:"max_idle_connections"`
	} `mapstructure:"db"`
	Secret struct {
		JwtKey       string `mapstructure:"jwt_key"`
		PassHashSalt string `mapstructure:"pass_hash_salt"`
	} `mapstructure:"secret"`
	Redis struct {
		Addr     string `mapstructure:"addr"`
		Password string `mapstructure:"password"`
		PoolSize int    `mapstructure:"pool_size"`
	} `mapstructure:"redis"`
}

func InitConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&Config); err != nil {
		panic(err)
	}
}
