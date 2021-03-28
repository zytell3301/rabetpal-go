package Controllers

import (
	"github.com/kataras/iris/v12"
	"rabetpal/Database/Cassandra/Models"
)

var Home = struct {
	Index func(ctx iris.Context)
}{
	Index: func(ctx iris.Context) {
		data := map[string]interface{}{"name": "Test", "lastname": "Test", "address": map[string]string{"address1": "addr1"},"mobile":"09372171814"}
		Models.NewUser(data,nil)
	},
}
