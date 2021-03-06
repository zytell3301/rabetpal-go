package Uuid

import (
	"fmt"
	"github.com/google/uuid"
)

var space uuid.UUID

func initSpace() {
	fmt.Println("Space uuid is being set")
	space = uuid.New()
	err := space.UnmarshalText([]byte(Configs.GetString("space")))
	switch err != nil {
	case true:
		panic("An error occurred while setting space uuid. Error: " + err.Error())
	}
	fmt.Println("Space uuid just got set successfully")
}

func GenerateV5(name string) string {
	return uuid.NewSHA1(space, []byte(name)).String()
}
