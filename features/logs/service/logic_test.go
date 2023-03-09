package service

import (
	"errors"
	"immersive-dashboard/features/logs"
	"immersive-dashboard/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	mock_data_log = logs.Core{
		ID:        1,
		MenteeID:  3,
		UserID:    4,
		StatusID:  3,
		Feedback:  "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
		CreatedAt: "2023-03-09",
	}
)

func TestAdd(t *testing.T) {
	repo := new(mocks.LogData)

	t.Run("Success Insert", func(t *testing.T) {
		repo.On("AddLogData", mock_data_log).Return(nil).Once()

		srv := New(repo)
		err := srv.AddLogSrv(mock_data_log)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed when func Insert return error", func(t *testing.T) {
		repo.On("AddLogData", mock_data_log).Return(errors.New("error insert data")).Once()

		srv := New(repo)
		err := srv.AddLogSrv(mock_data_log)
		assert.NotNil(t, err)
		assert.Equal(t, "error insert data", err.Error())
		repo.AssertExpectations(t)
	})

	t.Run("Failed validate", func(t *testing.T) {
		inputData := logs.Core{
			ID:       1,
			MenteeID: 4,
			UserID:   2,
			StatusID: 1,
		}
		srv := New(repo)
		err := srv.AddLogSrv(inputData)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}
