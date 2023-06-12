package notes

import (
	"NoteApp/initializer"

	"github.com/gin-gonic/gin"
)

func Apis(r *gin.Engine) {
	repository := InitRepository(initializer.Db)
	noteHandler := InitHandler(repository)

	r.POST("/notes", noteHandler.CreateNoteHandler)
	r.GET("/notes", noteHandler.ViewNotesHandler)
	r.DELETE("/notes", noteHandler.DeleteNoteHandler)
}
