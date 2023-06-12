package notes

import (
	"NoteApp/initializer"

	"github.com/gin-gonic/gin"
)

func Apis(r *gin.Engine) {
	repository := InitRepository(initializer.Db)
	noteHandler := InitHandler(repository)

	r.POST("/note", noteHandler.CreateNoteHandler)
	r.GET("/note", noteHandler.ViewNotesHandler)
	r.DELETE("/note", noteHandler.DeleteNoteHandler)
}
