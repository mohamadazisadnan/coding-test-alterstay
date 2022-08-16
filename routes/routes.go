package routes

import (
	"test/alterstay/factory"

	"github.com/labstack/echo/v4"
)

func New(presenter factory.Presenter) *echo.Echo {
	e := echo.New()

	e.POST("/articles", presenter.ArticlePresenter.InsertData)
	e.GET("/articles", presenter.ArticlePresenter.GetAllData)
	return e
}
