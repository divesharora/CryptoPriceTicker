package models

import (
	"github.com/jinzhu/gorm"
)

type Users struct {
	gorm.Model
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
}

