package sessions

import (
	"gorm.io/gorm"
)

func InitRepository(client *gorm.DB) Repository {
	return NewRepository(client)
}

func InitSession() Session{
	return NewSession()
}

func InitCookie() Cookie{
	return NewCookie()
}