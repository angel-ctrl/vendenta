package database

import (
	"github.com/vendenta/models"
)

func CreateUserExample(a models.UserExample) models.UserExample {

	DB.Create(&models.UserExample{Name: "dog3", Last_name: "dog4", EmailsUWU: []models.EmailExample{{Emails: "toy3"}, {Emails: "toy4"}}})
	//DB.Create(&a)

	return a
}
