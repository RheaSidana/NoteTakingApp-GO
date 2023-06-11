package main

import (
	"NoteApp/initializer"
	"NoteApp/src/model"
)

func init() {
	initializer.LoadEnvVariables()
	initializer.ConnectToDB()
}

func main() {
	initializer.Db.AutoMigrate(&model.User{})
	initializer.Db.AutoMigrate(&model.Session{})
	initializer.Db.AutoMigrate(&model.Note{})
}
