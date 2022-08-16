package data

import (
	"test/alterstay/features/categories"

	"gorm.io/gorm"
)

type Categories struct {
	gorm.Model
	Name string `json:"name" form:"name"`
}

func (data *Categories) toCore() categories.Core {
	return categories.Core{
		ID:   int(data.ID),
		Name: data.Name,
	}
}

func toCoreList(data []Categories) []categories.Core {
	result := []categories.Core{}
	for key := range data {
		result = append(result, data[key].toCore())
	}
	return result
}
