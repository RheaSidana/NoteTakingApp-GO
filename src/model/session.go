package model

import "gorm.io/gorm"

type Session struct {
	gorm.Model
	SessionID string `json:"sessionID" binding:"required"`
	UserId    uint   `json:"userID" binding:"required"`
}
