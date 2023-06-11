package users

import (
	"NoteApp/src/model"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Repository interface {
	Create(user model.User) (model.User, error)
	Find(user model.User) (model.User, error)
}

type repository struct{
	client *gorm.DB
}

func NewRepository(client *gorm.DB) Repository {
	return &repository{client: client}
}

func (r *repository) Create(user model.User) (model.User, error) {
	encryptPassword,err := bcrypt.GenerateFromPassword(
		[]byte(user.Password), 
		bcrypt.DefaultCost)
	if err != nil {
		return model.User{},errors.New("encryption unsuccessful")
	}

	user.Password = string(encryptPassword);
	result := r.client.Create(&user)
	
	if result.Error != nil {
		return model.User{}, result.Error
	}

	return user, nil
}

func (r *repository) Find(user model.User) (model.User, error) {
	password := user.Password
	
	res := r.client.Where("users.email=?", user.Email).Find(&user)
	if res.Error != nil {
		return model.User{}, res.Error
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return model.User{},errors.New("incorrect credentials")
	}

	return user, nil
}
