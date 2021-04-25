package Hash

import "golang.org/x/crypto/bcrypt"

func GenerateHash(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), Configs.GetInt("hash_round"))
	return string(hash)
}
