package users

import (
	"NoteApp/initializer"

	"github.com/gin-gonic/gin"
)

func Apis(r *gin.Engine) {
	repository := InitRepository(initializer.Db)
	userHandler := InitHandler(repository)

	r.POST("/signup", userHandler.SignUpUserHandler)
	r.POST("/login", userHandler.LoginUserHandler)
	r.POST("/logout", userHandler.LogoutUserHandler)
}
