package users

import (
	"NoteApp/src/model"
	mockSessions "NoteApp/src/modules/sessions/mocks"
	"NoteApp/src/modules/users/mocks"
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

func TestSignUpUserHandlerWhenEmptyUser(t *testing.T) {
	repo := new(mocks.Repository)
	handler := Handler{repository: repo}
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/signup", handler.SignUpUserHandler)
	newUser := model.User{}
	b, _ := json.Marshal(newUser)
	body := bytes.NewBuffer(b)
	req, _ := http.NewRequest(http.MethodPost, "/signup", body)
	respR := httptest.NewRecorder()
	expectedUser := newUser
	repo.On("Create", newUser).Return(model.User{}, errors.New("Empty User JSON."))

	actualUser, _ := repo.Create(newUser)
	r.ServeHTTP(respR, req)

	assert.Equal(t, respR.Code, http.StatusBadRequest)
	assert.Equal(t, expectedUser, actualUser)
}

func TestSignUpUserHandlerWhenUnableToCreateUser(t *testing.T) {
	repo := new(mocks.Repository)
	handler := Handler{repository: repo}
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/signup", handler.SignUpUserHandler)
	newUser := model.User{
		Name:     "test",
		Email:    "test@example.com",
		Password: "Test@3_5",
	}
	b, _ := json.Marshal(newUser)
	body := bytes.NewBuffer(b)
	req, _ := http.NewRequest(http.MethodPost, "/signup", body)
	respR := httptest.NewRecorder()
	expectedUser := newUser
	expectedUser.ID = 1
	repo.On("Create", newUser).Return(model.User{}, errors.New("Error while creating user"))

	actualUser, _ := repo.Create(newUser)
	r.ServeHTTP(respR, req)

	assert.Equal(t, respR.Code, http.StatusInternalServerError)
	assert.NotEqual(t, expectedUser, actualUser)
}

func TestSignUpUserHandler(t *testing.T) {
	repo := new(mocks.Repository)
	handler := Handler{repository: repo}
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/signup", handler.SignUpUserHandler)
	newUser := model.User{
		Name:     "test",
		Email:    "test@example.com",
		Password: "Test@3_5",
	}
	b, _ := json.Marshal(newUser)
	body := bytes.NewBuffer(b)
	req, _ := http.NewRequest(http.MethodPost, "/signup", body)
	respR := httptest.NewRecorder()
	expectedUser := newUser
	expectedUser.ID = 1
	repo.On("Create", newUser).Return(expectedUser, nil)

	actualUser, _ := repo.Create(newUser)
	r.ServeHTTP(respR, req)

	assert.Equal(t, respR.Code, http.StatusOK)
	assert.Equal(t, expectedUser, actualUser)
}

func TestLoginUserHandlerWhenEmptyUser(t *testing.T) {
	repo := new(mocks.Repository)
	handler := Handler{repository: repo}
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/login", handler.LoginUserHandler)
	newUser := model.User{}
	b, _ := json.Marshal(newUser)
	body := bytes.NewBuffer(b)
	req, _ := http.NewRequest(http.MethodPost, "/login", body)
	respR := httptest.NewRecorder()
	expectedUser := newUser
	repo.On("Find", newUser).Return(model.User{}, errors.New("Empty User JSON."))

	actualUser, _ := repo.Find(newUser)
	r.ServeHTTP(respR, req)

	assert.Equal(t, respR.Code, http.StatusBadRequest)
	assert.Equal(t, expectedUser, actualUser)
}

func TestLoginUserHandlerWhenUnableToCreateUser(t *testing.T) {
	repo := new(mocks.Repository)
	handler := Handler{repository: repo}
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/login", handler.LoginUserHandler)
	newUser := model.User{
		Name:     "test",
		Email:    "test@example.com",
		Password: "Test@3_5",
	}
	b, _ := json.Marshal(newUser)
	body := bytes.NewBuffer(b)
	req, _ := http.NewRequest(http.MethodPost, "/login", body)
	respR := httptest.NewRecorder()
	expectedUser := newUser
	expectedUser.ID = 1
	repo.On("Find", newUser).Return(model.User{}, errors.New("Error while creating user"))

	actualUser, _ := repo.Find(newUser)
	r.ServeHTTP(respR, req)

	assert.Equal(t, respR.Code, http.StatusUnauthorized)
	assert.NotEqual(t, expectedUser, actualUser)
}

func TestLoginUserHandlerWhenUnableToCreateSession(t *testing.T) {
	session := new(mockSessions.Session)
	repo := new(mocks.Repository)
	handler := Handler{repository: repo, session: session}
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/login", handler.LoginUserHandler)
	newUser := model.User{
		Email:    "test@example.com",
		Password: "Test@3_5",
	}
	b, _ := json.Marshal(newUser)
	body := bytes.NewBuffer(b)
	req, _ := http.NewRequest(http.MethodPost, "/login", body)
	respR := httptest.NewRecorder()
	expectedUser := newUser
	expectedUser.ID = 1
	repo.On("Find", newUser).Return(expectedUser, nil)
	session.On("GenerateSessionID", newUser, expectedUser.ID).Return("", errors.New("something went wrong"))

	actualUser, _ := repo.Find(newUser)
	actualSession, err := session.GenerateSessionID(newUser, expectedUser.ID)
	r.ServeHTTP(respR, req)

	assert.Equal(t, respR.Code, http.StatusInternalServerError)
	assert.Equal(t, expectedUser, actualUser)
	assert.Equal(t, "something went wrong", err.Error())
	assert.Equal(t, "", actualSession)
}

func TestLoginUserHandler(t *testing.T) {
	cookie := new(mockSessions.Cookie)
	session := new(mockSessions.Session)
	repo := new(mocks.Repository)
	handler := Handler{repository: repo, session: session, cookie: cookie}
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/login", handler.LoginUserHandler)
	newUser := model.User{
		Email:    "test@example.com",
		Password: "Test@3_5",
	}
	b, _ := json.Marshal(newUser)
	body := bytes.NewBuffer(b)
	req, _ := http.NewRequest(http.MethodPost, "/login", body)
	respR := httptest.NewRecorder()
	expectedUser := newUser
	expectedUser.ID = 1
	repo.On("Find", newUser).Return(expectedUser, nil)
	expectedSession := "abc$tRO09"
	session.On("GenerateSessionID", newUser, expectedUser.ID).Return(expectedSession, nil)
	cookie.On("SetCookie", mock.AnythingOfType("*gin.Context"), expectedSession).Return()

	actualUser, _ := repo.Find(newUser)
	actualSession, err := session.GenerateSessionID(newUser, expectedUser.ID)
	r.ServeHTTP(respR, req)

	assert.Equal(t, respR.Code, http.StatusOK)
	assert.Equal(t, expectedUser, actualUser)
	assert.Equal(t, nil, err)
	assert.Equal(t, expectedSession, actualSession)
}

func TestLogoutUserHandlerWhenEmptyUser(t *testing.T) {
	repo := new(mocks.Repository)
	handler := Handler{repository: repo}
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/logout", handler.LogoutUserHandler)
	newUser := UserRequestWithSession{}
	b, _ := json.Marshal(newUser)
	body := bytes.NewBuffer(b)
	req, _ := http.NewRequest(http.MethodPost, "/logout", body)
	respR := httptest.NewRecorder()

	r.ServeHTTP(respR, req)

	assert.Equal(t, respR.Code, http.StatusBadRequest)
}

func TestLogoutUserHandlerWhenInvalidSession(t *testing.T) {
	session := new(mockSessions.Session)
	repo := new(mocks.Repository)
	handler := Handler{repository: repo, session: session}
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/logout", handler.LogoutUserHandler)
	newUser := UserRequestWithSession{
		SID: "sessionID",
	}
	b, _ := json.Marshal(newUser)
	body := bytes.NewBuffer(b)
	req, _ := http.NewRequest(http.MethodPost, "/logout", body)
	respR := httptest.NewRecorder()
	session.On("IsAuthenticate", "sessionID").Return(model.Session{}, errors.New("unauthorised access"))

	r.ServeHTTP(respR, req)
	actualSession, err := session.IsAuthenticate("sessionID")

	assert.Equal(t, respR.Code, http.StatusUnauthorized)
	assert.Equal(t, model.Session{}, actualSession)
	assert.Equal(t, "unauthorised access", err.Error())
}

func TestLogoutUserHandlerWhenUnableToDeleteSessionID(t *testing.T) {
	session := new(mockSessions.Session)
	repo := new(mocks.Repository)
	handler := Handler{repository: repo, session: session}
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/logout", handler.LogoutUserHandler)
	newUser := UserRequestWithSession{
		SID: "sessionID",
	}
	b, _ := json.Marshal(newUser)
	body := bytes.NewBuffer(b)
	req, _ := http.NewRequest(http.MethodPost, "/logout", body)
	respR := httptest.NewRecorder()
	expectedSession := model.Session{
		SessionID: "sessionID",
		UserId: uint(1),
	}
	session.On("IsAuthenticate", "sessionID").Return(expectedSession, nil)
	session.On("DeleteSessionID", "sessionID").Return(model.Session{}, errors.New("unable to delete session"))

	r.ServeHTTP(respR, req)
	actualSession, err := session.IsAuthenticate("sessionID")

	assert.Equal(t, respR.Code, http.StatusInternalServerError)
	assert.Equal(t, expectedSession, actualSession)
	assert.Equal(t, nil, err)

	actualSession, err = session.DeleteSessionID("sessionID")
	assert.Equal(t, model.Session{}, actualSession)
	assert.Equal(t, "unable to delete session", err.Error())
}

func TestLogoutUserHandler(t *testing.T) {
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
	r.POST("/logout", handler.LogoutUserHandler)
	newUser := UserRequestWithSession{
		SID: "sessionID",
	}
	b, _ := json.Marshal(newUser)
	body := bytes.NewBuffer(b)
	req, _ := http.NewRequest(http.MethodPost, "/logout", body)
	respR := httptest.NewRecorder()
	expectedSession := model.Session{
		SessionID: "sessionID",
		UserId: uint(1),
	}
	session.On("IsAuthenticate", "sessionID").Return(expectedSession, nil)
	session.On("DeleteSessionID", "sessionID").Return(expectedSession, nil)
	cookie.On("DeleteCookie",  mock.AnythingOfType("*gin.Context")).Return()
	
	r.ServeHTTP(respR, req)
	assert.Equal(t, respR.Code, http.StatusOK)

	actualSession, err := session.IsAuthenticate("sessionID")
	
	assert.Equal(t, expectedSession, actualSession)
	assert.Equal(t, nil, err)
	
	actualSession, err = session.DeleteSessionID("sessionID")
	assert.Equal(t, expectedSession, actualSession)
	assert.Equal(t, nil, err)
}