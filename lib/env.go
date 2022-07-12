package lib

import (
	"fmt"

	"github.com/spf13/viper"
)

type Env struct {
	ServerPort string `mapstructure:"SERVER_PORT"`
}

var env Env

func GetEnv() *Env {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		panic("cannot read cofiguration")
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		fmt.Errorf("unable to decode into struct, %v", err)
	}
	return &env
}
