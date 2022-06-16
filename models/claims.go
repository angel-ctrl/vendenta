package models

import (
	jwt "github.com/dgrijalva/jwt-go"
)

type Claim struct {
	UserName string `json:"user"`
	ID    string `json:"_id"`
	jwt.StandardClaims
}
