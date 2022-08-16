package factory

import (
	"gorm.io/gorm"

	_categoryBusiness "test/alterstay/features/categories/business"
	_categoryData "test/alterstay/features/categories/data"
	_categoryPresentation "test/alterstay/features/categories/presentation"

	_articleBusiness "test/alterstay/features/articles/business"
	_articleData "test/alterstay/features/articles/data"
	_articlePresentation "test/alterstay/features/articles/presentation"
)

type Presenter struct {
	CategoryPresenter *_categoryPresentation.CategoryHandler
	ArticlePresenter  *_articlePresentation.ArticleHandler
}

func InitFactory(dbConn *gorm.DB) Presenter {
	categoryData := _categoryData.NewCategoryRepository(dbConn)
	categoryBusiness := _categoryBusiness.NewCategoryBusiness(categoryData)
	categoryPresentation := _categoryPresentation.NewCategoryHandler(categoryBusiness)

	articleData := _articleData.NewArticleRepository(dbConn)
	articleBuiness := _articleBusiness.NewArticleBusiness(articleData)
	articlePresentation := _articlePresentation.NewArticleHandler(articleBuiness)

	return Presenter{
		CategoryPresenter: categoryPresentation,
		ArticlePresenter:  articlePresentation,
	}
}
