package users

import (
	"NoteApp/src/model"
	"NoteApp/src/modules/sessions"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	repository Repository
	session    sessions.Session
	cookie     sessions.Cookie
}

// SignUp
func (h *Handler) SignUpUserHandler(c *gin.Context) {
	var newUser model.User
	c.BindJSON(&newUser)
	if newUser == (model.User{}) {
		c.JSON(400, ErrorResponse{Message: "Bad Request: Unable to add user."})
		return
	}

	user, err := h.repository.Create(newUser)
	if err != nil {
		c.JSON(500, ErrorResponse{Message: "Unable to add user."})
		return
	}

	c.JSON(200, UserResponse{Message: user.Name + " created successfully!!"})
}

// login
func (h *Handler) LoginUserHandler(c *gin.Context) {
	var user model.User
	c.ShouldBindJSON(&user)
	if user == (model.User{}) {
		c.JSON(400, ErrorResponse{Message: "Bad Request: Unable to add user."})
		return
	}

	userInDB, err := h.repository.Find(user)
	if err != nil {
		c.JSON(401, ErrorResponse{Message: "Unauthorised Access."})
		return
	}

	session, err := h.session.GenerateSessionID(user, userInDB.ID)
	if err != nil {
		c.JSON(500, ErrorResponse{Message: "Something went wrong."})
		return
	}

	h.cookie.SetCookie(c, session)

	c.JSON(200, UserRequestWithSession{SID: session})
}

// logout
func (h *Handler) LogoutUserHandler(c *gin.Context) {
	var userSession UserRequestWithSession
	c.ShouldBindJSON(&userSession)
	if userSession == (UserRequestWithSession{}) {
		c.JSON(400, ErrorResponse{Message: "Bad Request: Unable to add user."})
		return
	}

	session, err := h.session.IsAuthenticate(userSession.SID)
	if err != nil {
		c.JSON(401, ErrorResponse{Message: "Unauthorised Access."})
		return
	}

	session, err = h.session.DeleteSessionID(session.SessionID)
	if err != nil {
		c.JSON(500, ErrorResponse{Message: "Something went wrong."})
		return
	}

	h.cookie.DeleteCookie(c)

	c.JSON(200, UserResponse{Message: "User logged out successfully"})
}
