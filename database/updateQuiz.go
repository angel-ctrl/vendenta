package database

import (
	"fmt"

	"github.com/vendenta/models"
)

func Update_Quiz_Database(id string, modelUpdate *models.Quiz) models.Quiz {

	var QuizSearshed models.Quiz

	err := DB.Model(&QuizSearshed).Where("id = ?", id).Error

	if err != nil {
		fmt.Println(err)
	}

	DB.Unscoped().Delete(QuizSearshed)

	CreateQuiz(*modelUpdate)

	return QuizSearshed

}