package services

import "github.com/golang-jwt/jwt"

type Claims struct {
	jwt.StandardClaims
	ID      string `json:"id"`
	AdminId string `json:"adminId"`
}
