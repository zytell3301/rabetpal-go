package Hash

import "golang.org/x/crypto/bcrypt"

func GenerateHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), Configs.GetInt("hash_round"))
	return string(hash), err
}
