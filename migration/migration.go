package migration

import (
	"gorm.io/gorm"

	_migrateArticles "test/alterstay/features/articles/data"
	_migrateCategories "test/alterstay/features/categories/data"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&_migrateCategories.Categories{})
	db.AutoMigrate(&_migrateArticles.Articles{})
}
