package categories

type Core struct {
	ID   int
	Name string
}

type Business interface {
	GetAllData() (data []Core, err error)
}

type Data interface {
	GetAllData() (data []Core, err error)
}
