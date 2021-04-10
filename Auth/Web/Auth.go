package Web

import (
	"github.com/dgrijalva/jwt-go/v4"
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

func GetUser(ctx iris.Context) (jwt.MapClaims, error) {
	signedString := ctx.GetCookie("user")
	switch signedString == "" {
	case true:
		return nil, nil
	}
	user, error := Auth.DecodeJwt(signedString)
	switch error != nil || user == nil {
	case true:
		return nil, error
	}
	return user.Claims.(jwt.MapClaims), nil
}
