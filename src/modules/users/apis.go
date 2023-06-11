package users

import (
	"NoteApp/initializer"

	"github.com/gin-gonic/gin"
)

func Apis(r *gin.Engine) {
	repository := InitRepository(initializer.Db)
	userHandler := InitHandler(repository)

	r.POST("/signup", userHandler.CreateUserHandler)
	r.POST("/login", userHandler.FindUser)
}
