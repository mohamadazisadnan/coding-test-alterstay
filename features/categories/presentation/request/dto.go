package request

import "test/alterstay/features/categories"

type Categories struct {
	Name string `json:"name" form:"name"`
}

func ToCore(req Categories) categories.Core {
	return categories.Core{
		Name: req.Name,
	}
}
