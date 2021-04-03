package Auth

import "github.com/dgrijalva/jwt-go"

func KeyFunc(*jwt.Token) (interface{}, error){
	return "123bac12415baf148291fb1241561291",nil
}

func DecodeJwt(signedString string) (jwt.MapClaims,error){
	decodedJwt,err :=jwt.Parse(signedString,KeyFunc)
	return decodedJwt.Claims.(jwt.MapClaims),err
}