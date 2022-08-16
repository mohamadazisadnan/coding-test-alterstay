package business

import (
	"test/alterstay/features/articles"
)

type useCase struct {
	articleData articles.Data
}

func NewArticleBusiness(artlcData articles.Data) articles.Business {
	return &useCase{
		articleData: artlcData,
	}
}

func (uc *useCase) GetAllData() (data []articles.Core, err error) {
	data, err = uc.articleData.GetAllData()
	return data, err
}

func (uc *useCase) InsertData(insert articles.Core) (row int, err error) {
	row, err = uc.articleData.InsertData(insert)
	return row, err
}
