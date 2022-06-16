package database

import (
	"github.com/vendenta/models"
	"golang.org/x/crypto/bcrypt"
)

func EncriptarPass(pass string) (string, error) {

	costo := 8

	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	
	return string(bytes), err
}

func CreateUserProfile(a models.Account) models.Account {

	a.Password, _ = EncriptarPass(a.Password)

	DB.Create(&a)
	DB.Save(&a)

	return a 
}

/*type UserMarsh struct {
    *models.Account
}

func(a *UserMarsh) aaa(){
	a.Profile.
}*/
