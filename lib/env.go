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

	// Set Default env key values. Use key from file
	viper.SetDefault("SERVER_PORT", ":4000")

	err := viper.ReadInConfig()
	if err != nil {
		panic("cannot read cofiguration")
	}

	//To Watch for env file changes
	// viper.OnConfigChange(func(e fsnotify.Event) {
	// 	fmt.Println("Config file changed:", e.Name)
	// })
	// viper.WatchConfig()

	err = viper.Unmarshal(&env)
	if err != nil {
		fmt.Errorf("unable to decode into struct, %v", err)
	}
	return &env
}
