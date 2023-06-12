package notes

import (
	"NoteApp/src/model"

	"gorm.io/gorm"
)

type Repository interface {
	Create(note string, sessionID uint) (model.Note, error) 
	FindAll(sessionId uint) ([]model.Note, error)
	Delete(noteID uint) (model.Note, error)
}

type repository struct{
	client *gorm.DB
}

func NewRepository(client *gorm.DB) Repository {
	return &repository{client: client}
}

func (r *repository) Create(note string, sessionId uint) (model.Note, error) {
	noteObj := model.Note{
		Note: note,
		SessionId: sessionId,
	}

	result := r.client.Create(&noteObj)
	if result.Error != nil {
		return model.Note{}, result.Error
	}

	return noteObj, nil
}

func (r *repository) FindAll(sessionId uint) ([]model.Note, error){
	var notes []model.Note

	result := r.client.Where("session_id=?", sessionId).Find(&notes)
	if result.Error != nil {
		return []model.Note{}, result.Error
	}

	return notes, nil
}

func (r *repository) Delete(noteID uint) (model.Note, error){
	var note model.Note

	res := r.client.Where("id=?", noteID).Delete(&note)
	if res.Error != nil {
		return model.Note{}, res.Error
	}	

	return note, nil
}