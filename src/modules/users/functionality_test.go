package users

// import (
// 	"NoteApp/model"
// 	// "ScheduleJobs/modules/users/mocks"
// 	"errors"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// func TestGetUser(t *testing.T){
// 	repo := new(mocks.Repository)
// 	myUser := model.User{
// 		Name: "test",
// 		Email: "test@example.com",
// 		Password: "Test@3_5",
// 	}
// 	expectedUser := myUser
// 	expectedUser.ID = 1
// 	repo.On("Find", 1).Return(expectedUser,nil)

// 	user, _ := GetUser(1, repo)

// 	assert.Equal(t,expectedUser, user)
// }

// func TestGetUserNotFound(t *testing.T){
// 	repo := new(mocks.Repository)
// 	expectedUser := model.User{}
// 	expectedError := errors.New("Record not found")
// 	repo.On("Find", 1).Return(model.User{},expectedError)

// 	user, err := GetUser(1, repo)

// 	assert.Equal(t,expectedUser, user)
// 	assert.Error(t,expectedError, err)
// }

