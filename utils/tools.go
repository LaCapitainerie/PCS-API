package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword génère le hash bcrypt d'un mot de passe
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
