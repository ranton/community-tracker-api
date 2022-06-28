package hash

import (
	"golang.org/x/crypto/bcrypt"
)

const Cost int = 10

func Make(h string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(h), Cost)
	return string(hashed)
}

func Check(value string, hashedValue string) bool {
	hashed := Make(value)
	return hashed == hashedValue
}
