package data

import (
	"test/alterstay/features/categories"

	"gorm.io/gorm"
)

type mysqlCategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(conn *gorm.DB) categories.Data {
	return &mysqlCategoryRepository{
		db: conn,
	}
}

func (repo *mysqlCategoryRepository) GetAllData() (data []categories.Core, err error) {
	var getAllData []Categories
	tx := repo.db.Find(&getAllData)
	if tx.Error != nil {
		return []categories.Core{}, tx.Error
	}
	return toCoreList(getAllData), nil
}
