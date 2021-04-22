package main

import (
	"github.com/kataras/iris/v12"
	"rabetpal/Controllers"
	"rabetpal/Controllers/Users"
	_ "rabetpal/Database/Cassandra/Keyspaces"
	_ "rabetpal/Database/Uuid"
	_ "rabetpal/Database/Uuid/Generator"
)

func main() {
	r := iris.New()

	r.Get("/test",Controllers.HomeController.Index)
	r.Post("/register",Users.UsersController.Register)

	r.Listen(":4042")
}
