package routers

import (
	"errors"
	"strings"

	"github.com/vendenta/models"
	jwt "github.com/dgrijalva/jwt-go"
)

func ProcesoToken(tk string) (*models.Claim, bool, string, error) {

	miClave := []byte("hackernoob")

	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")

	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})

	if !tkn.Valid {
		return claims, false, string(""), errors.New("token invalido")
	}

	return claims, false, string(""), err
}
