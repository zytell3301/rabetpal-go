package main

import (
	"github.com/kataras/iris/v12"
	"rabetpal/Controllers"
	_ "rabetpal/Database/Cassandra/Keyspaces"
)

func main() {
	r := iris.New()

	r.Get("/test",Controllers.Home.Index)

	r.Listen(":4042")
}
