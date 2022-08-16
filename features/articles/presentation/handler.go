package presentation

import (
	"net/http"
	"test/alterstay/features/articles"
	"test/alterstay/features/articles/presentation/request"
	"test/alterstay/features/articles/presentation/response"

	"github.com/labstack/echo/v4"
)

type ArticleHandler struct {
	articleBusiness articles.Business
}

func NewArticleHandler(business articles.Business) *ArticleHandler {
	return &ArticleHandler{
		articleBusiness: business,
	}
}

func (h *ArticleHandler) GetAllData(c echo.Context) error {
	data, err := h.articleBusiness.GetAllData()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "Error",
			"message": "Failed to get all data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success to get all data",
		"data":    response.FromCoreList(data),
	})

}

func (h *ArticleHandler) InsertData(c echo.Context) error {
	var newArticle request.Articles
	errBind := c.Bind(&newArticle)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "Error",
			"message": "Failed to bind data, check your input",
		})
	}
	dataArticle := request.ToCore(newArticle)
	row, err := h.articleBusiness.InsertData(dataArticle)
	if row != 1 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "Error",
			"message": "Failed to insert data",
		})
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "Error",
		"message": "Success to insert data",
	})
}
