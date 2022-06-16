package database

import (
	"github.com/vendenta/models"
)

func SearchProfile(user string) (models.Account, bool) {

	var account models.Account

	if err := DB.Model(&account).Where("accounts.User = ?", string(user)).Preload("Profile").First(&account).Error; err != nil {
		return account, false
	}

	return account, true

}
