package notes

import (
	"NoteApp/src/modules/sessions"

	"gorm.io/gorm"
)

func InitRepository(client *gorm.DB) Repository {
	return NewRepository(client)
}

func InitHandler(userRepository Repository) Handler {
	return Handler{
		repository: userRepository,
		session: sessions.InitSession(),
		cookie: sessions.InitCookie(),
	}
}
