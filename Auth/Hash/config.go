package Hash

import (
	"fmt"
	"github.com/spf13/viper"
)

var Configs *viper.Viper

func init() {
	fmt.Println("Hash configs are being read")
	Configs = viper.New()
	Configs.SetConfigType("yaml")
	Configs.SetConfigName("Hash")
	Configs.AddConfigPath("./Configs")
	err := Configs.ReadInConfig()
	switch err != nil {
	case true:
		panic("Unable to read config \"Hash\". Error: " + err.Error())
	}
	fmt.Println("Hash configs just got set successfully")
}
