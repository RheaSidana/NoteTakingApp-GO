package sessions

import (
	"NoteApp/src/model"

	"gorm.io/gorm"
)

type Repository interface {
	Create(session model.Session) (model.Session, error)
	Find(sessionID string) (model.Session, error)
	FindByUser(userID uint) (model.Session, error)
	Delete(sessionID string) (model.Session, error)
}

type repository struct {
	client *gorm.DB
}

func NewRepository(client *gorm.DB) Repository {
	return &repository{client: client}
}

func (r *repository) Create(session model.Session) (model.Session, error) {
	result := r.client.Create(&session)

	if result.Error != nil {
		return model.Session{}, result.Error
	}

	return session, nil
}

func (r *repository) Find(sessionID string) (model.Session, error) {
	var session model.Session
	res := r.client.Where("session_id=?", sessionID).Find(&session)
	if res.Error != nil {
		return model.Session{}, res.Error
	}

	return session, nil
}

func (r *repository) FindByUser(userID uint) (model.Session, error) {
	var session model.Session
	res := r.client.Where("user_id=?", userID).Find(&session)
	if res.Error != nil {
		return model.Session{}, res.Error
	}

	return session, nil
}

func (r *repository) Delete(sessionID string) (model.Session, error) {
	var session model.Session
	res := r.client.Where("session_id=?", sessionID).Delete(&session)
	if res.Error != nil {
		return model.Session{}, res.Error
	}

	return session, nil
}
