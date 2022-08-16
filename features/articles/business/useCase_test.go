package business

import (
	"errors"
	"test/alterstay/features/articles"
	"test/alterstay/features/categories"
	"test/alterstay/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAllData(t *testing.T) {
	repo := new(mocks.ArticleData)
	returnData := []articles.Core{{ID: 1, Title: "Sample article 1", Category: categories.Core{
		ID: 1, Name: "success-story",
	}}}

	t.Run("Success Get All", func(t *testing.T) {
		repo.On("GetAllData").Return(returnData, nil).Once()

		srv := NewArticleBusiness(repo)

		res, err := srv.GetAllData()
		assert.NoError(t, err)
		assert.Equal(t, returnData[0].ID, res[0].ID)
		repo.AssertExpectations(t)
	})

	t.Run("Error Get All", func(t *testing.T) {
		repo.On("GetAllData").Return(nil, errors.New("data not found")).Once()

		srv := NewArticleBusiness(repo)

		res, err := srv.GetAllData()
		assert.Error(t, err)
		assert.Nil(t, res)
		repo.AssertExpectations(t)
	})
}

func TestInsertData(t *testing.T) {
	repo := new(mocks.ArticleData)
	insertData := articles.Core{ID: 1, Title: "Simple article 1", Category: categories.Core{
		ID: 1,
	}}

	t.Run("Success Insert", func(t *testing.T) {
		repo.On("InsertData", mock.Anything).Return(1, nil).Once()
		srv := NewArticleBusiness(repo)

		res, err := srv.InsertData(insertData)
		assert.NoError(t, err)
		assert.Equal(t, 1, res)
		repo.AssertExpectations(t)
	})

}
