package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go" 
	"github.com/vendenta/models"
)

func GeneroJWT(t models.Account) (string, error){

	miClave := []byte("hackernoob")

	payload := jwt.MapClaims{
		"ID": t.ID,
		"exp": time.Now().Add(time.Minute * 60).Unix(), 
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(miClave)

	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil

}