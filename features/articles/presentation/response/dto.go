package response

import "test/alterstay/features/articles"

type Articles struct {
	ID       int        `json:"id"`
	Title    string     `json:"title"`
	Category Categories `json:"category"`
}

type Categories struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func FromCore(data articles.Core) Articles {
	return Articles{
		ID:    data.ID,
		Title: data.Title,
		Category: Categories{
			ID:   data.Category.ID,
			Name: data.Category.Name,
		},
	}
}

func FromCoreList(data []articles.Core) []Articles {
	result := []Articles{}
	for key := range data {
		result = append(result, FromCore(data[key]))
	}
	return result
}
