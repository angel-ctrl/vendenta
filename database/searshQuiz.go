package database

import (
	"fmt"

	"github.com/vendenta/models"
)

func SearshQuizDB(id string) models.Quiz{

	var QuizSearshed models.Quiz

	err := DB.Model(&QuizSearshed).Where("id = ?", id).Preload("Questions").Preload("Questions.Answer").Find(&QuizSearshed).Error

	if err != nil {
		fmt.Println(err)
	}

	return QuizSearshed

}
