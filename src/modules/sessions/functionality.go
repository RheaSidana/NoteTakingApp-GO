package sessions

import (
	"NoteApp/initializer"
	"NoteApp/src/model"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func GenerateSessionID(user model.User, id uint) (string, error) {
	today := time.Now().UTC().String()
	user.ID = id

	repository := InitRepository(initializer.Db)
	session, _ := repository.FindByUser(user.ID)
	if session != (model.Session{}) {
		DeleteSessionID(session.SessionID)
	}

	sessionID, err := bcrypt.GenerateFromPassword(
		[]byte(user.Email+user.Password+today), 
		bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("session encryption unsuccessful")
	}

	repository.Create(model.Session{
		SessionID: string(sessionID),
		UserId: user.ID,	
	})
	if err != nil {
		return "", errors.New("session failed to add to the table")
	}

	return string(sessionID), nil
}

func IsAuthenticate(sessionID string) (model.Session, error){
	repository := InitRepository(initializer.Db)

	session, err := repository.Find(sessionID)
	if err != nil {
		return model.Session{}, errors.New("encryption unsuccessful")
	}

	return session, nil
}

func DeleteSessionID(sessionID string) (model.Session, error){
	repository := InitRepository(initializer.Db)

	session, err := repository.Delete(sessionID)
	if err != nil {
		return model.Session{}, errors.New("encryption unsuccessful")
	}

	return session, nil
}
