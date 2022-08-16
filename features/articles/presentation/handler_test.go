package presentation

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"test/alterstay/features/articles"
	"test/alterstay/features/articles/presentation/request"
	"test/alterstay/features/articles/presentation/response"
	"test/alterstay/features/categories"
	"test/alterstay/mocks"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ArticlesResponseSuccess struct {
	Message string
	Data    []response.Articles
}

type ResponseGlobal struct {
	Message string
}

var (
	mock_data_user = request.Articles{
		Title:      "Sample article 1",
		CategoryID: 1,
	}
)

func TestGetAll(t *testing.T) {
	e := echo.New()
	usecase := new(mocks.ArticleBusiness)
	returnData := []articles.Core{{ID: 1, Title: "Sample article 1", Category: categories.Core{
		ID: 1, Name: "success-story",
	}}}

	t.Run("Success Get All", func(t *testing.T) {
		usecase.On("GetAllData", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(returnData, nil).Once()
		srv := NewArticleHandler(usecase)

		req := httptest.NewRequest(http.MethodGet, "/articles", nil)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/articles")
		// if you want to add param, add this code below
		// context.SetParamNames("id")
		// context.SetParamValues("1")

		responseData := ArticlesResponseSuccess{}
		// result :=
		if assert.NoError(t, srv.GetAllData(echoContext)) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			// fmt.Println("res", responseBody)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, returnData[0].Title, responseData.Data[0].Title)
		}
		usecase.AssertExpectations(t)
	})

	t.Run("Error Get All", func(t *testing.T) {
		usecase.On("GetAllData", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil, errors.New("Failed to get all data")).Once()

		srv := NewArticleHandler(usecase)

		req := httptest.NewRequest(http.MethodGet, "/articles", nil)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/articles")

		srv.GetAllData(echoContext)
		responseBody := rec.Body.String()
		var responseData ResponseGlobal
		err := json.Unmarshal([]byte(responseBody), &responseData)
		fmt.Println("res", responseData)
		if err != nil {
			assert.Error(t, err, "error")
		}
		// assert.Error(t, result)
		assert.Equal(t, "Failed to get all data", responseData.Message)
		// assert.Nil(t, res)
		usecase.AssertExpectations(t)
	})
}

func TestInsertData(t *testing.T) {
	reqBody, err := json.Marshal(mock_data_user)
	if err != nil {
		t.Error(t, err, "error")
	}

	e := echo.New()
	usecase := new(mocks.ArticleBusiness)
	// returnData := []users.Core{{ID: 1, Name: "alta", Email: "alta@mail.id", Password: "qwerty"}}

	t.Run("Success to insert data", func(t *testing.T) {
		usecase.On("InsertData", mock.Anything).Return(1, nil).Once()

		srv := NewArticleHandler(usecase)

		req := httptest.NewRequest(http.MethodPost, "/articles", bytes.NewBuffer(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/articles")

		responseData := ResponseGlobal{}

		if assert.NoError(t, srv.InsertData(echoContext)) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			// fmt.Println("res", responseBody)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, "Success to insert data", responseData.Message)
		}
		usecase.AssertExpectations(t)
	})

	t.Run("Failed Add User when bind error", func(t *testing.T) {

		var dataFail = map[string]int{
			"title": 1,
		}
		reqBodyFail, _ := json.Marshal(dataFail)
		srv := NewArticleHandler(usecase)

		req := httptest.NewRequest(http.MethodPost, "/articles", bytes.NewBuffer(reqBodyFail))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/articles")

		responseData := ResponseGlobal{}

		if assert.NoError(t, srv.InsertData(echoContext)) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			// fmt.Println("res", responseBody)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, "Failed to bind data, check your input", responseData.Message)
		}
		usecase.AssertExpectations(t)
	})

	t.Run("Failed Add User when insert data failed", func(t *testing.T) {
		usecase.On("InsertData", mock.Anything).Return(-1, errors.New("failed to insert data")).Once()
		// var dataFail = map[string]int{
		// 	"name": 1,
		// }
		// reqBodyFail, _ := json.Marshal(dataFail)
		srv := NewArticleHandler(usecase)

		req := httptest.NewRequest(http.MethodPost, "/articles", bytes.NewBuffer(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/articles")

		responseData := ResponseGlobal{}

		if assert.NoError(t, srv.InsertData(echoContext)) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			// fmt.Println("res", responseBody)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, "Failed to insert data", responseData.Message)
		}
		usecase.AssertExpectations(t)
	})

	t.Run("Failed Add User when insert data failed", func(t *testing.T) {
		usecase.On("InsertData", mock.Anything).Return(0, errors.New("failed to insert data")).Once()

		srv := NewArticleHandler(usecase)

		req := httptest.NewRequest(http.MethodPost, "/articles", bytes.NewBuffer(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/articles")

		responseData := ResponseGlobal{}

		if assert.NoError(t, srv.InsertData(echoContext)) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			// fmt.Println("res", responseBody)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, "Failed to insert data", responseData.Message)
		}
		usecase.AssertExpectations(t)
	})

}
