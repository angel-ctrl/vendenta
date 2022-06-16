package database

import (
	"errors"
	"github.com/vendenta/models"
	"golang.org/x/crypto/bcrypt"
)

func Loginatempt(u models.Account, password string) (bool, error) {

	passwordBytes := []byte(password)

	passwordBD := []byte(u.Password)

	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)

	if err != nil {
		return false, errors.New("contraseña incorrecta")  
	}

	return true, nil

}