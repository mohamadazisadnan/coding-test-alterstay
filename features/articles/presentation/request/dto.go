package request

import (
	"test/alterstay/features/articles"
	"test/alterstay/features/categories"
)

type Articles struct {
	Title      string `json:"title" form:"title"`
	CategoryID uint   `json:"category_id" form:"category_id"`
}

func ToCore(req Articles) articles.Core {
	return articles.Core{
		Title: req.Title,
		Category: categories.Core{
			ID: int(req.CategoryID),
		},
	}
}
