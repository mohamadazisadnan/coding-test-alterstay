package articles

import (
	"test/alterstay/features/categories"
)

type Core struct {
	ID       int
	Title    string
	Category categories.Core
}

type Business interface {
	InsertData(insert Core) (row int, err error)
	GetAllData() (data []Core, err error)
}

type Data interface {
	InsertData(insert Core) (row int, err error)
	GetAllData() (data []Core, err error)
}
