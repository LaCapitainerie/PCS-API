package models

import "github.com/dgrijalva/jwt-go"

// Claim est la structure spécifiant les données utilisées pour les tokens de connexion
type Claims struct {
	IdUser string `json:"idUser"`
	jwt.StandardClaims
}
