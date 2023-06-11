package main

import (
	"NoteApp/initializer"
	"NoteApp/src/modules/users"

	"github.com/gin-gonic/gin"
)

func init() {
	initializer.LoadEnvVariables()
	initializer.ConnectToDB()
}

func main() {
	r := gin.Default()

	users.Apis(r)

	r.Run()
}
