package service

import (
	"immersive-dashboard/features/status"
	"immersive-dashboard/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelectAll(t *testing.T) {
	repo := new(mocks.StatusData)
	returnData := []status.Core{
		{
			ID:   1,
			Name: "Immersive Back End 15",
		},
	}

	t.Run("Success Get All", func(t *testing.T) {
		repo.On("SelectAll").Return(returnData, nil).Once()

		srv := New(repo)
		response, err := srv.GetAll()
		assert.Nil(t, err)
		assert.Equal(t, returnData[0].Name, response[0].Name)
		repo.AssertExpectations(t)
	})
}
