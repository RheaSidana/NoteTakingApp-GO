package model

import "gorm.io/gorm"

type Note struct {
	gorm.Model
	Note      string `json:"note" binding:"required"`
	SessionId uint   `gorm:"references:Sessions.ID" json:"sessionID" binding:"required"`
}
