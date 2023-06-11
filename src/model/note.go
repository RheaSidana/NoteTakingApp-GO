package model

import "gorm.io/gorm"

type Note struct {
	gorm.Model
	Id   uint32 `json:"userID" binding:"required"`
	Note string `json:"sessionID" binding:"required"`
}
