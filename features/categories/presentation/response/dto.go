package response

import "test/alterstay/features/categories"

type Categories struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func FromCore(data categories.Core) Categories {
	return Categories{
		ID:   data.ID,
		Name: data.Name,
	}
}

func FromCoreList(data []categories.Core) []Categories {
	result := []Categories{}
	for key := range data {
		result = append(result, FromCore(data[key]))
	}
	return result
}
