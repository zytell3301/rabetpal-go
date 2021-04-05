package Controllers

import (
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/kataras/iris/v12"
	"rabetpal/Auth"
)

var HomeController = struct {
	Index func(ctx iris.Context)
}{
	Index: func(ctx iris.Context) {
		signedString, _ := Auth.EncodeJwt(jwt.MapClaims{"name": "Test", "lastname": "Test", "address": map[string]string{"address1": "addr1"}, "mobile": "09372171814"})
		fmt.Println(signedString)
		fmt.Println(Auth.DecodeJwt(signedString))
	},
}
