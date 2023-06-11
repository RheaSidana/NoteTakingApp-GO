package sessions

import (
	"NoteApp/initializer"
	"NoteApp/src/model"
	"errors"
	"fmt"

	// "fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func GenerateSessionID(user model.User) (string, error) {
	today := time.Now().UTC().String()

	sessionID, err := bcrypt.GenerateFromPassword(
		[]byte(user.Email+user.Password+today), 
		bcrypt.DefaultCost)
	fmt.Println("error: ", err)
	if err != nil {
		return "", errors.New("session encryption unsuccessful")
	}


	repository := InitRepository(initializer.Db)
	repository.Create(model.Session{
		SessionID: string(sessionID),
		UserId: user.ID,	
	})
	if err != nil {
		return "", errors.New("session failed to add to the table")
	}

	// fmt.Println("encrypt", string(encrypt))

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
