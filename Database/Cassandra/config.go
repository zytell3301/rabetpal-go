package Cassandra

import (
	"fmt"
	"github.com/spf13/viper"
)

var Configs *viper.Viper

func init() {
	fmt.Println("Config is being set")
	Configs = viper.New()
	Configs.AddConfigPath("./Configs")
	Configs.SetConfigName("Cassandra")
	Configs.SetConfigType("yaml")
	err := Configs.ReadInConfig()
	switch err != nil {
	case true:
		panic("Unable to read config \"Cassandra\". Error: " + err.Error())
	}
}
