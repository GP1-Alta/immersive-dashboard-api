package service

import (
	"errors"
	"immersive-dashboard/features/mentees"
	"immersive-dashboard/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	mock_data_mentee = mentees.Core{
		ID:              1,
		Name:            "Naruto Uzumaki",
		Address:         "Konohagakure",
		HomeAddress:     "Konohagakure",
		Email:           "naruto@alta.com",
		Sex:             "Male",
		Telegram:        "@narutouzumaki",
		Phone:           "08123456789",
		Discord:         "naruto#5060",
		StatusID:        1,
		ClassID:         1,
		EmergencyName:   "Minato Namikaze",
		EmergencyPhone:  "0123456789",
		EmergencyStatus: "Orang Tua",
		Category:        "Informatics",
		Major:           "Teknik Informatika",
		Institution:     "Institut Teknologi Konoha",
		CreatedAt:       time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt:       time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
	}
)

func TestDelete(t *testing.T) {
	repo := new(mocks.MenteeData)

	t.Run("Success Delete", func(t *testing.T) {
		id := uint(1)
		repo.On("Delete", mock_data_mentee, id).Return(nil).Once()

		srv := New(repo)
		err := srv.Delete(mock_data_mentee, id)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed when func Delete return error", func(t *testing.T) {
		id := uint(1)
		repo.On("Delete", mock_data_mentee, id).Return(errors.New("error delete data class")).Once()

		srv := New(repo)
		err := srv.Delete(mock_data_mentee, id)
		assert.NotNil(t, err)
		assert.Equal(t, "error delete data class", err.Error())
		repo.AssertExpectations(t)
	})
}

func TestInsert(t *testing.T) {
	repo := new(mocks.MenteeData)

	t.Run("Success Insert", func(t *testing.T) {
		repo.On("Insert", mock_data_mentee).Return(nil).Once()

		srv := New(repo)
		err := srv.Create(mock_data_mentee)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed when func Insert return error", func(t *testing.T) {
		repo.On("Insert", mock_data_mentee).Return(errors.New("error insert data")).Once()

		srv := New(repo)
		err := srv.Create(mock_data_mentee)
		assert.NotNil(t, err)
		assert.Equal(t, "error insert data", err.Error())
		repo.AssertExpectations(t)
	})

	t.Run("Failed validate", func(t *testing.T) {
		inputData := mentees.Core{
			ID:              1,
			Name:            "Naruto Uzumaki",
			Address:         "Konohagakure",
			HomeAddress:     "Konohagakure",
			Email:           "naruto@alta.com",
			Sex:             "Male",
			Telegram:        "@narutouzumaki",
			Phone:           "08123456789",
			Discord:         "naruto#5060",
			StatusID:        1,
			ClassID:         1,
			EmergencyPhone:  "0123456789",
			EmergencyStatus: "Orang Tua",
			Category:        "Informatics",
			Major:           "Teknik Informatika",
			Institution:     "Institut Teknologi Konoha",
		}
		srv := New(repo)
		err := srv.Create(inputData)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestSelect(t *testing.T) {
	repo := new(mocks.MenteeData)
	idMentee := uint(1)

	t.Run("Success Select One", func(t *testing.T) {
		repo.On("Select", mock.AnythingOfType("uint")).Return(mock_data_mentee, nil).Once()

		srv := New(repo)
		response, err := srv.Get(idMentee)
		assert.Nil(t, err)
		assert.Equal(t, mock_data_mentee.Name, response.Name)
		repo.AssertExpectations(t)
	})
}

func TestSelectAll(t *testing.T) {
	repo := new(mocks.MenteeData)
	returnData := []mentees.Core{
		{
			ID:       1,
			Name:     "Naruto Uzumaki",
			Class:    "Immersive Back End 15",
			Status:   "Join Class",
			Category: "Informatics",
		},
	}
	page := 1
	limit := 10
	offset := (page - 1) * limit
	class := "Immersive Back End 15"
	status := "Join Class"
	category := "Informatics"
	name := "Naruto"

	t.Run("Success Get All", func(t *testing.T) {
		repo.On("SelectAll", limit, offset, class, status, category, name).Return(returnData, nil).Once()

		srv := New(repo)
		response, err := srv.GetAll(page, class, status, category, name)
		assert.Nil(t, err)
		assert.Equal(t, returnData[0].Name, response[0].Name)
		repo.AssertExpectations(t)
	})
}
