package presentation

import (
	"net/http"
	"test/alterstay/features/categories"
	"test/alterstay/features/categories/presentation/response"

	"github.com/labstack/echo/v4"
)

type CategoryHandler struct {
	categoryBusiness categories.Business
}

func NewCategoryHandler(business categories.Business) *CategoryHandler {
	return &CategoryHandler{
		categoryBusiness: business,
	}
}

func (h *CategoryHandler) GetAllData(c echo.Context) error {
	data, err := h.categoryBusiness.GetAllData()
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
