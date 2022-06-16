package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Account struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();"primaryKey""`
	User      string    
	Password  string    `json:"Password" binding:"required"`
	Active    bool      `json:"Active" binding:"required"`

	Profile   ProfileAccount `gorm:"ForeignKey:ProfileID" json:"Profile"`
	ProfileID uuid.UUID `gorm:"type:uuid"`

	CreatedAt int64          `gorm:"autoCreateTime:milli" json:"created_at"`
	UpdatedAt int64          `gorm:"autoUpdateTime:milli" json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type ProfileAccount struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();"primaryKey""`
	FirstName string    `json:"firstname" binding:"required"`
	LastName  string    `json:"lastname" binding:"required"`
	Age       int       `json:"age" binding:"required"`
}