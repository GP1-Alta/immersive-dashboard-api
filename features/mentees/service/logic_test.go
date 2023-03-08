package service

import (
	"errors"
	"immersive-dashboard/features/mentees"
	"immersive-dashboard/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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
