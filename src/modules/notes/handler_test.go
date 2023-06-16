package notes

import (
	"NoteApp/src/model"
	"NoteApp/src/modules/notes/mocks"
	mockSessions "NoteApp/src/modules/sessions/mocks"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateNoteHandlerWhenEmptyRequestBody(t *testing.T) {
	seesion := new(mockSessions.Session)
	repo := new(mocks.Repository)
	handler := Handler{repository: repo, session: seesion}
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/notes", handler.CreateNoteHandler)
	newNote := CreateNoteRequestBody{}
	b, _ := json.Marshal(newNote)
	body := bytes.NewBuffer(b)
	req, _ := http.NewRequest(http.MethodPost, "/notes", body)
	respR := httptest.NewRecorder()
	expectedNote := model.Note{}
	sessionID := uint(1)
	repo.On("Create", "note", sessionID).Return(expectedNote, errors.New("Empty Note JSON."))

	actualNote, _ := repo.Create("note", sessionID)
	r.ServeHTTP(respR, req)

	assert.Equal(t, respR.Code, http.StatusBadRequest)
	assert.Equal(t, expectedNote, actualNote)
}


func TestCreateNoteHandlerWhenUnathorisedAccessThroughSession(t *testing.T) {
	session := new(mockSessions.Session)
	repo := new(mocks.Repository)
	handler := Handler{repository: repo, session: session}
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/notes", handler.CreateNoteHandler)
	newNote := CreateNoteRequestBody{}
	b, _ := json.Marshal(newNote)
	body := bytes.NewBuffer(b)
	req, _ := http.NewRequest(http.MethodPost, "/notes", body)
	respR := httptest.NewRecorder()
	expectedNote := model.Note{}
	sessionID := uint(1)
	repo.On("Create", "note", sessionID).Return(expectedNote, errors.New("Empty Note JSON."))
	session.On("IsAuthenticate", "sessionID").Return(model.Session{}, errors.New("unauthorised access"))

	actualNote, _ := repo.Create("note", sessionID)
	actualSession, err := session.IsAuthenticate("sessionID")
	r.ServeHTTP(respR, req)

	assert.Equal(t, respR.Code, http.StatusBadRequest)
	assert.Equal(t, expectedNote, actualNote)
	assert.Equal(t, model.Session{}, actualSession)
	assert.Equal(t, "unauthorised access", err.Error())
}

func TestCreateNoteHandlerWhenUnathorisedAccessThroughCookieCompare(t *testing.T) {
	cookie := new(mockSessions.Cookie)
	session := new(mockSessions.Session)
	repo := new(mocks.Repository)
	handler := Handler{
		repository: repo, 
		session: session, 
		cookie: cookie,
	}
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/notes", handler.CreateNoteHandler)
	newNote := CreateNoteRequestBody{}
	b, _ := json.Marshal(newNote)
	body := bytes.NewBuffer(b)
	req, _ := http.NewRequest(http.MethodPost, "/notes", body)
	respR := httptest.NewRecorder()
	expectedNote := model.Note{}
	expectedSession := model.Session{
		SessionID: "sessionID",
		UserId: uint(1),
	}
	sessionID := uint(1)
	repo.On("Create", "note", sessionID).Return(expectedNote, errors.New("Empty Note JSON."))
	session.On("IsAuthenticate", "sessionID").Return(expectedSession,nil)
	cookie.On("GetCookie", mock.AnythingOfType("*gin.Context")).Return("", errors.New("unauthorised access"))

	actualNote, _ := repo.Create("note", sessionID)
	assert.Equal(t, expectedNote, actualNote)

	actualSession, err := session.IsAuthenticate("sessionID")
	assert.Equal(t, expectedSession, actualSession)
	assert.Equal(t, nil, err)
	
	var Context *gin.Context
	
	actualCookie, err := cookie.GetCookie(Context)
	assert.Equal(t, "", actualCookie)
	assert.Equal(t, "unauthorised access", err.Error())
	
	r.ServeHTTP(respR, req)
	assert.Equal(t, respR.Code, http.StatusBadRequest)
}

func TestCreateNoteHandlerWhenUnableToCreateNote(t *testing.T) {
	cookie := new(mockSessions.Cookie)
	session := new(mockSessions.Session)
	repo := new(mocks.Repository)
	handler := Handler{
		repository: repo, 
		session: session, 
		cookie: cookie,
	}
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/notes", handler.CreateNoteHandler)
	requestBody := CreateNoteRequestBody{
		Sid : "sessionID",
		Note: "note",
	}
	b, _ := json.Marshal(requestBody)
	body := bytes.NewBuffer(b)
	req, _ := http.NewRequest(http.MethodPost, "/notes", body)
	respR := httptest.NewRecorder()
	expectedSession := model.Session{
		SessionID: requestBody.Sid,
		UserId: uint(1),
	}
	sessionID_userID := uint(1)
	expectedSession.ID = sessionID_userID
	expectedNote := model.Note{
		Note: requestBody.Note,
		SessionId: sessionID_userID,
	}
	repo.On("Create", "note", sessionID_userID).Return(model.Note{}, errors.New("Empty Note JSON."))
	session.On("IsAuthenticate", requestBody.Sid).Return(expectedSession,nil)
	cookie.On("GetCookie", mock.AnythingOfType("*gin.Context")).Return(requestBody.Sid, nil)

	actualNote, _ := repo.Create("note", sessionID_userID)
	assert.NotEqual(t, expectedNote, actualNote)

	actualSession, err := session.IsAuthenticate(requestBody.Sid)
	assert.Equal(t, expectedSession, actualSession)
	assert.Equal(t, nil, err)
	
	var Context *gin.Context
	actualCookie, err := cookie.GetCookie(Context)
	assert.Equal(t, requestBody.Sid, actualCookie)
	assert.Equal(t, nil, err)
	
	
	r.ServeHTTP(respR, req)
	assert.Equal(t, respR.Code, http.StatusInternalServerError)
}

func TestCreateNoteHandler(t *testing.T) {
	cookie := new(mockSessions.Cookie)
	session := new(mockSessions.Session)
	repo := new(mocks.Repository)
	handler := Handler{
		repository: repo, 
		session: session, 
		cookie: cookie,
	}
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/notes", handler.CreateNoteHandler)
	requestBody := CreateNoteRequestBody{
		Sid : "sessionID",
		Note: "note",
	}
	b, _ := json.Marshal(requestBody)
	body := bytes.NewBuffer(b)
	req, _ := http.NewRequest(http.MethodPost, "/notes", body)
	respR := httptest.NewRecorder()
	expectedSession := model.Session{
		SessionID: requestBody.Sid,
		UserId: uint(1),
	}
	sessionID_userID := uint(1)
	expectedSession.ID = sessionID_userID
	expectedNote := model.Note{
		Note: requestBody.Note,
		SessionId: sessionID_userID,
	}
	repo.On("Create", "note", sessionID_userID).Return(expectedNote, nil)
	session.On("IsAuthenticate", requestBody.Sid).Return(expectedSession,nil)
	cookie.On("GetCookie", mock.AnythingOfType("*gin.Context")).Return(requestBody.Sid, nil)

	actualNote, _ := repo.Create("note", sessionID_userID)
	assert.Equal(t, expectedNote, actualNote)

	actualSession, err := session.IsAuthenticate(requestBody.Sid)
	assert.Equal(t, expectedSession, actualSession)
	assert.Equal(t, nil, err)
	
	var Context *gin.Context
	actualCookie, err := cookie.GetCookie(Context)
	assert.Equal(t, requestBody.Sid, actualCookie)
	assert.Equal(t, nil, err)
	
	
	r.ServeHTTP(respR, req)
	assert.Equal(t, respR.Code, http.StatusOK)
}

func TestViewNoteHandlerWhenEmptyRequestBody(t *testing.T) {
	seesion := new(mockSessions.Session)
	repo := new(mocks.Repository)
	handler := Handler{repository: repo, session: seesion}
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/notes", handler.ViewNotesHandler)
	newNote := ViewNoteRequestBody{}
	b, _ := json.Marshal(newNote)
	body := bytes.NewBuffer(b)
	req, _ := http.NewRequest(http.MethodGet, "/notes", body)
	respR := httptest.NewRecorder()

	r.ServeHTTP(respR, req)

	assert.Equal(t, respR.Code, http.StatusBadRequest)
}

func TestViewNoteHandlerWhenUnableToFindAnyNote(t *testing.T) {
	cookie := new(mockSessions.Cookie)
	session := new(mockSessions.Session)
	repo := new(mocks.Repository)
	handler := Handler{
		repository: repo, 
		session: session,
		cookie: cookie,
	}
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/notes", handler.ViewNotesHandler)
	requestBody := ViewNoteRequestBody{
		Sid : "sessionID",
	}
	b, _ := json.Marshal(requestBody)
	body := bytes.NewBuffer(b)
	req, _ := http.NewRequest(http.MethodGet, "/notes", body)
	respR := httptest.NewRecorder()
	expectedSession := model.Session{
		SessionID: requestBody.Sid,
		UserId: uint(1),
	}
	sessionID_userID := uint(1)
	expectedSession.ID = sessionID_userID
	cookie.On("GetCookie", mock.AnythingOfType("*gin.Context")).Return(requestBody.Sid, nil)
	session.On("IsAuthenticate", requestBody.Sid).Return(expectedSession,nil)
	repo.On("FindAll", sessionID_userID).Return([]model.Note{}, errors.New("Empty Note JSON."))

	r.ServeHTTP(respR, req)

	
	actualSession, err := session.IsAuthenticate(requestBody.Sid)
	assert.Equal(t, expectedSession, actualSession)
	assert.Equal(t, nil, err)
	
	var Context *gin.Context
	actualCookie, err := cookie.GetCookie(Context)
	assert.Equal(t, requestBody.Sid, actualCookie)
	assert.Equal(t, nil, err)
	
	notes, err := repo.FindAll(sessionID_userID)
	assert.Equal(t, "Empty Note JSON.", err.Error())
	assert.Equal(t, []model.Note{}, notes)
	
	assert.Equal(t, respR.Code, http.StatusInternalServerError)
}

func TestViewNoteHandler(t *testing.T) {
	cookie := new(mockSessions.Cookie)
	session := new(mockSessions.Session)
	repo := new(mocks.Repository)
	handler := Handler{
		repository: repo, 
		session: session,
		cookie: cookie,
	}
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/notes", handler.ViewNotesHandler)
	requestBody := ViewNoteRequestBody{
		Sid : "sessionID",
	}
	b, _ := json.Marshal(requestBody)
	body := bytes.NewBuffer(b)
	req, _ := http.NewRequest(http.MethodGet, "/notes", body)
	respR := httptest.NewRecorder()
	expectedSession := model.Session{
		SessionID: requestBody.Sid,
		UserId: uint(1),
	}
	sessionID_userID := uint(1)
	expectedSession.ID = sessionID_userID
	cookie.On("GetCookie", mock.AnythingOfType("*gin.Context")).Return(requestBody.Sid, nil)
	session.On("IsAuthenticate", requestBody.Sid).Return(expectedSession,nil)
	expectedNote := model.Note{
		Note: "note",
		SessionId: sessionID_userID,
	}
	repo.On("FindAll", sessionID_userID).Return([]model.Note{expectedNote}, nil)

	r.ServeHTTP(respR, req)

	
	actualSession, err := session.IsAuthenticate(requestBody.Sid)
	assert.Equal(t, expectedSession, actualSession)
	assert.Equal(t, nil, err)
	
	var Context *gin.Context
	actualCookie, err := cookie.GetCookie(Context)
	assert.Equal(t, requestBody.Sid, actualCookie)
	assert.Equal(t, nil, err)
	
	notes, err := repo.FindAll(sessionID_userID)
	assert.Equal(t, nil, err)
	assert.Equal(t, []model.Note{expectedNote}, notes)
	
	assert.Equal(t, respR.Code, http.StatusOK)
}

func TestDeleteNoteHandlerWhenEmptyRequestBody(t *testing.T) {
	seesion := new(mockSessions.Session)
	repo := new(mocks.Repository)
	handler := Handler{repository: repo, session: seesion}
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.DELETE("/notes", handler.DeleteNoteHandler)
	newNote := DeleteNoteRequestBody{}
	b, _ := json.Marshal(newNote)
	body := bytes.NewBuffer(b)
	req, _ := http.NewRequest(http.MethodDelete, "/notes", body)
	respR := httptest.NewRecorder()

	r.ServeHTTP(respR, req)

	assert.Equal(t, respR.Code, http.StatusBadRequest)
}

func TestDeleteNoteHandlerWhenUnableToDeleteNote(t *testing.T) {
	cookie := new(mockSessions.Cookie)
	session := new(mockSessions.Session)
	repo := new(mocks.Repository)
	handler := Handler{
		repository: repo, 
		session: session,
		cookie: cookie,
	}
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.DELETE("/notes", handler.DeleteNoteHandler)
	requestBody := DeleteNoteRequestBody{
		Sid: "sessionID",
		NoteID: uint(1),
	}
	b, _ := json.Marshal(requestBody)
	body := bytes.NewBuffer(b)
	req, _ := http.NewRequest(http.MethodDelete, "/notes", body)
	respR := httptest.NewRecorder()
	expectedSession := model.Session{
		SessionID: requestBody.Sid,
		UserId: uint(1),
	}
	sessionID_userID := uint(1)
	expectedSession.ID = sessionID_userID
	cookie.On("GetCookie", mock.AnythingOfType("*gin.Context")).Return(requestBody.Sid, nil)
	session.On("IsAuthenticate", requestBody.Sid).Return(expectedSession,nil)
	repo.On("Delete", requestBody.NoteID).Return(model.Note{}, errors.New("unable to delete note"))

	r.ServeHTTP(respR, req)

	assert.Equal(t, respR.Code, http.StatusInternalServerError)

	actualSession, err := session.IsAuthenticate(requestBody.Sid)
	assert.Equal(t, expectedSession, actualSession)
	assert.Equal(t, nil, err)
	
	var Context *gin.Context
	actualCookie, err := cookie.GetCookie(Context)
	assert.Equal(t, requestBody.Sid, actualCookie)
	assert.Equal(t, nil, err)

	note, err := repo.Delete(requestBody.NoteID)
	assert.Equal(t, model.Note{}, note)
	assert.Equal(t, "unable to delete note", err.Error())
}

func TestDeleteNoteHandler(t *testing.T) {
	cookie := new(mockSessions.Cookie)
	session := new(mockSessions.Session)
	repo := new(mocks.Repository)
	handler := Handler{
		repository: repo, 
		session: session,
		cookie: cookie,
	}
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.DELETE("/notes", handler.DeleteNoteHandler)
	requestBody := DeleteNoteRequestBody{
		Sid: "sessionID",
		NoteID: uint(1),
	}
	b, _ := json.Marshal(requestBody)
	body := bytes.NewBuffer(b)
	req, _ := http.NewRequest(http.MethodDelete, "/notes", body)
	respR := httptest.NewRecorder()
	expectedSession := model.Session{
		SessionID: requestBody.Sid,
		UserId: uint(1),
	}
	sessionID_userID := uint(1)
	expectedSession.ID = sessionID_userID
	cookie.On("GetCookie", mock.AnythingOfType("*gin.Context")).Return(requestBody.Sid, nil)
	session.On("IsAuthenticate", requestBody.Sid).Return(expectedSession,nil)
	expectedNote := model.Note{
		Note: "note",
		SessionId: sessionID_userID,
	}
	repo.On("Delete", requestBody.NoteID).Return(expectedNote, nil)

	r.ServeHTTP(respR, req)

	assert.Equal(t, respR.Code, http.StatusOK)

	actualSession, err := session.IsAuthenticate(requestBody.Sid)
	assert.Equal(t, expectedSession, actualSession)
	assert.Equal(t, nil, err)
	
	var Context *gin.Context
	actualCookie, err := cookie.GetCookie(Context)
	assert.Equal(t, requestBody.Sid, actualCookie)
	assert.Equal(t, nil, err)

	note, err := repo.Delete(requestBody.NoteID)
	assert.Equal(t, expectedNote, note)
	assert.Equal(t, nil, err)
}
