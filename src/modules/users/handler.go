package users

import (
	"NoteApp/src/model"
	"NoteApp/src/modules/sessions"

	"github.com/gin-gonic/gin"
)

type Handler struct{
	repository Repository
}

// SignUp
func (h *Handler) CreateUserHandler(c *gin.Context){
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

//login
func (h *Handler) FindUser(c *gin.Context){
	var user model.User
	c.ShouldBindJSON(&user)
	if user == (model.User{}) {
		c.JSON(400, ErrorResponse{Message: "Bad Request: Unable to add user."})
		return
	}
	
	_, err := h.repository.Find(user)
	if err != nil {
		c.JSON(401, ErrorResponse{Message: "Unauthorised Access."})
		return
	}
	
	session, err := sessions.GenerateSessionID(user)
	if err != nil {
		c.JSON(500, ErrorResponse{Message: "Something went wrong."})
		return
	}

	c.JSON(200, UserResponseWithSession{SID: session})
}