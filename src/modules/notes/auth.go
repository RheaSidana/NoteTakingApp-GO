package notes

import (
	"NoteApp/src/model"
	"NoteApp/src/modules/sessions"
	"errors"

	"github.com/gin-gonic/gin"
)

func Authenticate(c *gin.Context, sessionID string) (model.Session, error){
	session, err := sessions.IsAuthenticate(sessionID)
	if err != nil {
		c.JSON(401, ErrorResponse{Message: "Unauthorised Access."})
		return model.Session{}, errors.New("unauthorised access")
	}
	
	cookie, err := sessions.GetCookie(c)
	if cookie == "" || cookie != session.SessionID && err != nil {
		c.JSON(401, ErrorResponse{Message: "Unauthorised Access."})
		return model.Session{}, errors.New("unauthorised access")
	}

	return session, nil
}

func AuthenticateUsingCookie(c *gin.Context) (model.Session, error){
	cookie, err := sessions.GetCookie(c)
	if cookie == "" || err != nil {
		c.JSON(401, ErrorResponse{Message: "Unauthorised Access."})
		return model.Session{}, errors.New("unauthorised access")
	}

	session, err := sessions.IsAuthenticate(cookie)
	if err != nil {
		c.JSON(401, ErrorResponse{Message: "Unauthorised Access."})
		return model.Session{}, errors.New("unauthorised access")
	}

	return session, nil
 }