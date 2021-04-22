package Generator

import (
	"fmt"
	"github.com/google/uuid"
	"rabetpal/Database/Uuid"
)

var space uuid.UUID

func init() {
	space = uuid.New()
	err := space.UnmarshalText([]byte(Uuid.Configs.GetString("space")))
	switch err != nil {
	case true:
		panic("An error occurred while setting space uuid. Error: " + err.Error())
	}
	fmt.Println("Space uuid just got set successfully")
}
