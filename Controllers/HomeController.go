package Controllers

import (
	"net/http"
	"rabetpal/Database/Cassandra/Models"
)

var Home = struct {
	Index func(response http.ResponseWriter, request *http.Request)
}{
	Index: func(response http.ResponseWriter, request *http.Request) {
		data := map[string]interface{}{"name": "Arshiya", "lastname": "Kiani", "address": map[string]string{"address1": "addr1"}}
		Models.NewUser(data)
	},
}
