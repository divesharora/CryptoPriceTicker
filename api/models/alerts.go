package models

import "github.com/jinzhu/gorm"

type Alert struct {
	gorm.Model
	Email  string  `json:"email" gorm:"not null"`
	Symbol string  `json:"symbol" gorm:"not null"`
	Price  string `json:"price" gorm:"not null"`
	Status int     `json:"status" gorm:"not null default:0"`
}

//Status 0: Not triggered, 1: Triggered ,2:Deleted
