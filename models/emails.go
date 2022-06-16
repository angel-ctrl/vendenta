package models

import (
	"gorm.io/gorm"
  )

type UserExample struct {
	gorm.Model
	//Id string `gorm:"primaryKey"`
	Name      string `json:"name" binding:"required"`
	Last_name string `json:"last_name" binding:"required"`
	EmailsUWUID uint
	EmailsUWU []EmailExample `gorm:"foreignkey:OwnerID"`
}

type EmailExample struct {
	gorm.Model
	Emails string 
	OwnerID   int
}
