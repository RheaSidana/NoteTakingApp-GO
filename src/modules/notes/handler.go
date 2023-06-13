package notes

import (
	"NoteApp/src/modules/sessions"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	repository Repository
	session sessions.Session
	cookie sessions.Cookie
}

func (h *Handler) CreateNoteHandler(c *gin.Context) {
	var sessionNote CreateNoteRequestBody
	c.ShouldBindJSON(&sessionNote)
	if sessionNote == (CreateNoteRequestBody{}) {
		c.JSON(400, ErrorResponse{Message: "Bad Request: Unable to add note."})
		return
	}

	session, err := Authenticate(c, sessionNote.Sid, h.session, h.cookie)
	if err != nil {
		return
	}

	noteCreated, err := h.repository.Create(sessionNote.Note, session.ID)
	if err != nil {
		c.JSON(500, ErrorResponse{Message: "Unable to add note."})
		return
	}

	c.JSON(200, CreateNoteResponseBody{ID: noteCreated.ID})
}

func (h *Handler) ViewNotesHandler(c *gin.Context) {
	var sessionID ViewNoteRequestBody
	c.ShouldBindJSON(&sessionID)
	if sessionID == (ViewNoteRequestBody{}) {
		c.JSON(400, ErrorResponse{Message: "Bad Request: Unable to add note."})
		return
	}

	// test if the user has access or not
	userSession, err := AuthenticateUsingCookie(c,h.session, h.cookie)
	if err != nil {
		return
	}

	// return the session object of the session id provided in the request body
	session, err := h.session.IsAuthenticate(sessionID.Sid)
	if err != nil {
		c.JSON(500, ErrorResponse{Message: "Invalid Session ID."})
		return
	}
	// check if the session requested by user is the logged in user's session or not
	if userSession.UserId != session.UserId {
		c.JSON(500, ErrorResponse{Message: "Invalid Session ID."})
		return
	}

	notes, err := h.repository.FindAll(session.ID)
	if err != nil {
		c.JSON(500, ErrorResponse{Message: "Unable to find note."})
		return
	}

	c.JSON(200, GetNoteResponse{Notes: notes})
}

func (h *Handler) DeleteNoteHandler(c *gin.Context) {
	var sessionNote DeleteNoteRequestBody
	c.ShouldBindJSON(&sessionNote)
	if sessionNote == (DeleteNoteRequestBody{}) {
		c.JSON(400, ErrorResponse{Message: "Bad Request: Unable to add note."})
		return
	}

	// test if the user has access or not
	userSession, err := AuthenticateUsingCookie(c, h.session, h.cookie)
	if err != nil {
		return
	}

	// return the session object of the session id provided in the request body
	session, err := h.session.IsAuthenticate(sessionNote.Sid)
	if err != nil {
		c.JSON(500, ErrorResponse{Message: "Invalid Session ID."})
		return
	}
	// check if the session requested by user is the logged in user's session or not
	if userSession.UserId != session.UserId {
		c.JSON(500, ErrorResponse{Message: "Invalid Session ID."})
		return
	}

	_, err = h.repository.Delete(sessionNote.NoteID)
	if err != nil {
		c.JSON(500, ErrorResponse{Message: "Unable to find note."})
		return
	}

	c.JSON(200, SuccessResponse{
		Message: "Note deleted successfully",
	})
}