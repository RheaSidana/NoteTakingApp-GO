package notes

import "NoteApp/src/model"

type ErrorResponse struct {
	Message string
}

type SuccessResponse struct {
	Message string
}

type CreateNoteRequestBody struct {
	Sid  string `json:"sid"`
	Note string `json:"note"`
}

type CreateNoteResponseBody struct {
	ID uint
}

type ViewNoteRequestBody struct {
	Sid  string `json:"sid"`
}

type GetNoteResponse struct {
	Notes []model.Note
}

type DeleteNoteRequestBody struct {
	Sid string `json:"sid"`
	NoteID uint `json:"id"`
}