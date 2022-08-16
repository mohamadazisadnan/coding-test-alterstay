package data

import (
	"gorm.io/gorm"

	"test/alterstay/features/articles"
	"test/alterstay/features/categories"
	_categories "test/alterstay/features/categories/data"
)

type Articles struct {
	gorm.Model
	Title        string                 `json:"title" form:"title"`
	CategoriesID uint                   `json:"categorys_id" form:"categorys_id"`
	Categories   _categories.Categories `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (data *Articles) toCore() articles.Core {
	return articles.Core{
		ID:    int(data.ID),
		Title: data.Title,
		Category: categories.Core{
			ID:   int(data.Categories.ID),
			Name: data.Categories.Name,
		},
	}
}

func toCoreList(data []Articles) []articles.Core {
	result := []articles.Core{}
	for key := range data {
		result = append(result, data[key].toCore())
	}
	return result
}

func fromCore(core articles.Core) Articles {
	return Articles{
		Title:        core.Title,
		CategoriesID: uint(core.Category.ID),
	}
}
