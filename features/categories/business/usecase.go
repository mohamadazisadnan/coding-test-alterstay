package business

import (
	"test/alterstay/features/categories"
)

type userCase struct {
	categoryData categories.Data
}

func NewCategoryBusiness(ctgData categories.Data) categories.Business {
	return &userCase{
		categoryData: ctgData,
	}
}

func (uc *userCase) GetAllData() (data []categories.Core, err error) {
	data, err = uc.categoryData.GetAllData()
	return data, err
}
