package util

import "golang.org/x/crypto/bcrypt"

func CompareHash(hashedPassword, password, secretKey string) (bool, error) {
	combined := password + secretKey
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(combined))
	return err == nil, err
}
