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
