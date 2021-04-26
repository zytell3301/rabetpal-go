package Uuid

import (
	"fmt"
	"github.com/spf13/viper"
)

var Configs *viper.Viper

func init() {
	fmt.Println("Uuid Configs are being read")
	Configs = viper.New()
	Configs.AddConfigPath("./Configs")
	Configs.SetConfigName("Uuid")
	Configs.SetConfigType("yaml")
	err := Configs.ReadInConfig()
	switch err != nil {
	case true:
		panic("Unable to read config \"Uuid\". Error: " + err.Error())
	}
	fmt.Println("Config \"Uuid\" just got read successfully")
	initSpace()
}
