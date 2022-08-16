package data

import (
	"test/alterstay/config"
	"test/alterstay/features/articles"
	"test/alterstay/features/categories"
	"testing"

	"github.com/stretchr/testify/assert"
)

var DbConn = config.InitDBTest()

func TestGetAllData(t *testing.T) {
	DbConn.Migrator().DropTable(&Articles{})
	DbConn.AutoMigrate(&Articles{})
	mockArticle := articles.Core{ID: 2, Title: "Sample article 1", Category: categories.Core{
		ID: 1, Name: "success-story",
	}}
	dataInput := fromCore(mockArticle)
	DbConn.Save(&dataInput)
	repo := NewArticleRepository(DbConn)

	t.Run("Test Get All Data User", func(t *testing.T) {
		dataResult, err := repo.GetAllData()
		assert.Nil(t, err)
		//assert.Equal(t, 1, row)
		assert.Equal(t, mockArticle.Title, dataResult[0].Title)
	})
}

func TestInsertData(t *testing.T) {
	DbConn.Migrator().DropTable(&Articles{})
	DbConn.AutoMigrate(&Articles{})

	repo := NewArticleRepository(DbConn)

	t.Run("Test insert data", func(t *testing.T) {
		mockUser := articles.Core{Title: "Simple article 1", Category: categories.Core{
			ID: 1,
		}}
		row, err := repo.InsertData(mockUser)
		assert.Nil(t, err)
		assert.Equal(t, 1, row)
	})
}

func TestInsertDataFailed(t *testing.T) {
	DbConn.Migrator().DropTable(&Articles{})
	// DbConn.AutoMigrate(&User{})

	repo := NewArticleRepository(DbConn)

	t.Run("Test Create User", func(t *testing.T) {
		mockUser := articles.Core{Title: "Simple article 1", Category: categories.Core{
			ID: 1,
		}}
		row, err := repo.InsertData(mockUser)
		assert.NotNil(t, err)
		assert.Equal(t, 0, row)
	})
}
