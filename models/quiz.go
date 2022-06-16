package models

import (
	"gorm.io/gorm"
)

type Quiz struct {
	gorm.Model
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Max_point   int    `json:"max_point" binding:"required"`
	Time        int    `json:"time" binding:"required"`

	Questions []*Questions `gorm:"foreignkey:QuizID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Questions struct {
	gorm.Model
	Question string `json:"question" binding:"required"`
	Point    int    `json:"point" binding:"required"`
	Ordering int    `json:"ordering" binding:"required"`

	QuizID int `gorm:"foreignKey:ID"`

	Answer []*Answerds `gorm:"foreignkey:QuestionID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Answerds struct {
	gorm.Model
	Answer     string `json:"answer" binding:"Answer"`
	Correct    bool   `json:"correct" binding:"Correct"`
	QuestionID int    `gorm:"foreignKey:ID"`
}
