package service

import (
	"errors"
	"immersive-dashboard/features/classes"
	"immersive-dashboard/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	mock_data_class = classes.Core{
		ID:        1,
		Name:      "Immersive Back End 15",
		UserID:    1,
		StartDate: time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
		EndDate:   time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
		CreatedAt: time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt: time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
	}
)

func TestDelete(t *testing.T) {
	repo := new(mocks.ClassData)

	t.Run("Success Delete Class", func(t *testing.T) {
		id := uint(1)
		repo.On("Delete", mock_data_class, id).Return(nil).Once()

		srv := New(repo)
		err := srv.Delete(mock_data_class, id)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed when func Delete return error", func(t *testing.T) {
		id := uint(1)
		repo.On("Delete", mock_data_class, id).Return(errors.New("error delete data class")).Once()

		srv := New(repo)
		err := srv.Delete(mock_data_class, id)
		assert.NotNil(t, err)
		assert.Equal(t, "error delete data class", err.Error())
		repo.AssertExpectations(t)
	})
}

func TestInsert(t *testing.T) {
	repo := new(mocks.ClassData)

	t.Run("Success Insert Class", func(t *testing.T) {
		repo.On("Insert", mock_data_class).Return(nil).Once()

		srv := New(repo)
		err := srv.Create(mock_data_class)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed when func Insert return error", func(t *testing.T) {
		repo.On("Insert", mock_data_class).Return(errors.New("error insert data")).Once()

		srv := New(repo)
		err := srv.Create(mock_data_class)
		assert.NotNil(t, err)
		assert.Equal(t, "error insert data", err.Error())
		repo.AssertExpectations(t)
	})

	t.Run("Failed validate", func(t *testing.T) {
		inputData := classes.Core{
			ID:        1,
			Name:      "Immersive Back End 15",
			StartDate: time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
			EndDate:   time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
		}
		srv := New(repo)
		err := srv.Create(inputData)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestListAll(t *testing.T) {
	repo := new(mocks.ClassData)
	returnData := []classes.Core{
		{ID: 1, Name: "Immersive Back End 15"},
	}

	t.Run("Success Get All Users", func(t *testing.T) {
		repo.On("ListAll").Return(returnData, nil).Once()

		srv := New(repo)
		response, err := srv.List()
		assert.Nil(t, err)
		assert.Equal(t, returnData[0].Name, response[0].Name)
		repo.AssertExpectations(t)
	})
}

func TestSelectAll(t *testing.T) {
	repo := new(mocks.ClassData)
	returnData := []classes.Core{
		{
			ID:           1,
			Name:         "Immersive Back End 15",
			Mentor:       "Iruka Sensei",
			StartDateStr: "2023-01-23",
			EndDateStr:   "2023-03-23",
		},
	}

	t.Run("Success Get All", func(t *testing.T) {
		repo.On("SelectAll", mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("string")).Return(returnData, nil).Once()

		srv := New(repo)
		response, err := srv.GetAll(1, "Back End")
		assert.Nil(t, err)
		assert.Equal(t, returnData[0].Name, response[0].Name)
		repo.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	repo := new(mocks.ClassData)
	idClass := uint(1)

	t.Run("Success Update", func(t *testing.T) {
		repo.On("Update", mock_data_class, mock.AnythingOfType("uint")).Return(nil).Once()

		srv := New(repo)
		err := srv.Edit(mock_data_class, idClass)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed when func Update return error", func(t *testing.T) {
		repo.On("Update", mock_data_class, mock.AnythingOfType("uint")).Return(errors.New("error update data")).Once()

		srv := New(repo)
		err := srv.Edit(mock_data_class, idClass)
		assert.NotNil(t, err)
		assert.Equal(t, "error update data", err.Error())
		repo.AssertExpectations(t)
	})
}
