package service

import (
	"errors"
	"immersive-dashboard/features/users"
	"immersive-dashboard/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	mock_data_user = users.Core{
		Id:       1,
		Name:     "Gintama",
		Email:    "gintama@go.com",
		Password: "gin123",
		Role:     "Default",
		Team:     "Mentor",
		Status:   "Active",
	}
)

var (
	mock_data_get_profile = users.Core{
		Id:     1,
		Name:   "Gintama",
		Email:  "gintama@go.com",
		Role:   "Default",
		Team:   "Mentor",
		Status: "Active",
	}
)

func TestDelete(t *testing.T) {
	repo := new(mocks.UserData)

	t.Run("Success Delete", func(t *testing.T) {
		id := 1
		repo.On("DeleteData", id).Return(nil).Once()

		srv := New(repo)
		err := srv.DeleteSrv(id)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed when func Delete return error", func(t *testing.T) {
		id := 1
		repo.On("DeleteData", id).Return(errors.New("error delete data user")).Once()

		srv := New(repo)
		err := srv.DeleteSrv(id)
		assert.NotNil(t, err)
		assert.Equal(t, "error delete data user", err.Error())
		repo.AssertExpectations(t)
	})
}

func TestRegister(t *testing.T) {
	repo := new(mocks.UserData)

	t.Run("Success Register", func(t *testing.T) {
		repo.On("RegisterData", mock.Anything).Return(nil).Once()

		srv := New(repo)
		err := srv.RegisterSrv(mock_data_user)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed when func Insert return error", func(t *testing.T) {
		repo.On("RegisterData", mock.Anything).Return(errors.New("error insert data")).Once()

		srv := New(repo)
		err := srv.RegisterSrv(mock_data_user)
		assert.NotNil(t, err)
		assert.Equal(t, "error insert data", err.Error())
		repo.AssertExpectations(t)
	})

	t.Run("Failed validate", func(t *testing.T) {
		inputData := users.Core{
			Id:       3,
			Email:    "romeo@mail.com",
			Password: "asdasds",
			Role:     "Default",
			Team:     "Mentor",
			Status:   "Acitve",
		}
		srv := New(repo)
		err := srv.RegisterSrv(inputData)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

}

func TestUpdate(t *testing.T) {
	repo := new(mocks.UserData)
	id := 1

	t.Run("Success Update", func(t *testing.T) {
		repo.On("UpdateUserData", mock.Anything, mock.Anything).Return(nil).Once()

		srv := New(repo)
		err := srv.UpdateUserSrv(id, mock_data_user)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Update", func(t *testing.T) {
		repo.On("UpdateUserData", mock.Anything, mock.Anything).Return(errors.New("error")).Once()

		srv := New(repo)
		err := srv.UpdateUserSrv(id, mock_data_user)
		assert.NotNil(t, err)
		assert.Equal(t, "error", err.Error())
		repo.AssertExpectations(t)
	})
}

func TestGetMentor(t *testing.T) {
	repo := new(mocks.UserData)
	returnData := []users.Core{
		{
			Id:   1,
			Name: "Lionel Messi",
		},
	}

	t.Run("Success Get All", func(t *testing.T) {
		repo.On("GetMentorData").Return(returnData, nil).Once()

		srv := New(repo)
		response, err := srv.GetMentorSrv()
		assert.Nil(t, err)
		assert.Equal(t, returnData[0].Name, response[0].Name)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Get All", func(t *testing.T) {
		repo.On("GetMentorData").Return(nil, errors.New("error data")).Once()
		srv := New(repo)
		response, err := srv.GetMentorSrv()
		assert.NotNil(t, err)
		assert.Equal(t, response, []users.Core([]users.Core(nil)))
		repo.AssertExpectations(t)
	})
}

func TestGetUsers(t *testing.T) {
	repo := new(mocks.UserData)
	returnData := []users.Core{
		{
			Id:     1,
			Name:   "Lionel Messi",
			Email:  "messi@mail.com",
			Role:   "Default",
			Team:   "Mentor",
			Status: "Active",
		},
	}
	page := 1
	keyword := "Messi"

	t.Run("Success Get All", func(t *testing.T) {
		repo.On("GetUser", page, keyword).Return(returnData, nil).Once()

		srv := New(repo)
		response, err := srv.GetUser(page, keyword)
		assert.Nil(t, err)
		assert.Equal(t, returnData[0].Name, response[0].Name)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Get All", func(t *testing.T) {
		repo.On("GetUser", page, keyword).Return(nil, errors.New("error data")).Once()
		srv := New(repo)
		response, err := srv.GetUser(page, keyword)
		assert.NotNil(t, err)
		assert.Equal(t, response, []users.Core([]users.Core(nil)))
		repo.AssertExpectations(t)
	})
}

func TestProfileData(t *testing.T) {
	repo := new(mocks.UserData)
	idUser := 1

	t.Run("Success Select One", func(t *testing.T) {
		repo.On("ProfileData", mock.AnythingOfType("int")).Return(mock_data_get_profile, nil).Once()

		srv := New(repo)
		response, err := srv.ProfileSrv(idUser)
		assert.Nil(t, err)
		assert.Equal(t, mock_data_get_profile.Name, response.Name)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Get Profile", func(t *testing.T) {
		repo.On("ProfileData", mock.AnythingOfType("int")).Return(users.Core{}, errors.New("error data")).Once()
		srv := New(repo)
		response, err := srv.ProfileSrv(idUser)
		assert.NotNil(t, err)
		assert.Equal(t, response, users.Core{})
		repo.AssertExpectations(t)
	})
}
