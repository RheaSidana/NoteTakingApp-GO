package users

import (
	"gorm.io/gorm"
)

func InitRepository(client *gorm.DB) Repository {
	return NewRepository(client)
}

func InitHandler(userRepository Repository) Handler {
	return Handler{repository: userRepository}
}
