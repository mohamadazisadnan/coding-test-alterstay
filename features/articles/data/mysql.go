package data

import (
	"fmt"
	"test/alterstay/features/articles"

	"gorm.io/gorm"
)

type mysqlArticleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(conn *gorm.DB) articles.Data {
	return &mysqlArticleRepository{
		db: conn,
	}
}

func (repo *mysqlArticleRepository) GetAllData() (data []articles.Core, err error) {
	var getAllData []Articles
	tx := repo.db.Preload("Categories").Find(&getAllData)
	if tx.Error != nil {
		return []articles.Core{}, tx.Error
	}
	return toCoreList(getAllData), nil
}

func (repo *mysqlArticleRepository) InsertData(insert articles.Core) (row int, err error) {
	insertData := fromCore(insert)
	tx := repo.db.Create(&insertData)
	if tx.Error != nil {
		return 0, tx.Error
	}
	if tx.RowsAffected != 1 {
		return 0, fmt.Errorf("failed to insert data")
	}
	return int(tx.RowsAffected), nil
}
