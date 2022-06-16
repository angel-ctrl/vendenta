package database

import (
	//"github.com/google/uuid"

	"github.com/vendenta/models"
)

func CreateQuiz(a models.Quiz) models.Quiz {

	DB.Create(&a)

	return a
}
