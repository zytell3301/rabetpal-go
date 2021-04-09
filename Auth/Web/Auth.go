package Web

import (
	"github.com/kataras/iris/v12"
	"net/http"
	"rabetpal/Auth"
)

func LoginUser(user map[string]interface{}, ctx iris.Context) error {
	signedJwt, error := Auth.EncodeJwt(user)
	switch error != nil {
	case true:
		return error
	}
	ctx.SetCookie(&http.Cookie{
		Name:     "user",
		Value:    signedJwt,
		HttpOnly: true,
	})
	return nil
}
