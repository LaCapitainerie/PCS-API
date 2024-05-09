package utils

import (
	"PCS-API/models"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// HashPassword génère le hash bcrypt d'un mot de passe
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// CreateToken génère un token à utiliser, prends en paramètre l'id d'un utilisateur et renvoie le token et l'erreur
func CreateToken(id string) (string, error) {
	expirationTime := time.Now().Add(TokenExpirationTime)
	claims := &models.Claims{
		IdUser: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return token.SignedString(TokenKey)
}

// CheckPassword vérifie si le hash et le mot de passe correspondent bien
func CheckPassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// IsInArrayString vérifie si un chaine ne se trouve pas dans un tableau de chaine
func IsInArrayString(chaine string, tab []string) bool {
	for _, str := range tab {
		if str == chaine {
			return true
		}
	}
	return false
}
