package Auth

import (
	"github.com/dgrijalva/jwt-go/v4"
)

var KeyFunc = jwt.KnownKeyfunc(jwt.SigningMethodHS512, []byte("404E635266556A586E3272357538782F4125442A472D4B6150645367566B59703373367639792442264528482B4D6251655468576D5A7134743777217A25432A"))

func DecodeJwt(signedString string) (*jwt.Token, error) {
	decodedJwt, err := jwt.Parse(signedString, KeyFunc)
	return decodedJwt, err
}
