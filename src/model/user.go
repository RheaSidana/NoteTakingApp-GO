package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" binding:"required"`
	Email    string `gorm:"unique;not_null" json:"email" binding:"required"` 
	Password string `json:"password" binding:"required"`
}
