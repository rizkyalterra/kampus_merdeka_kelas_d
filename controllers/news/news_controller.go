package news

import (
	"kelasd/configs"
	"kelasd/models/news"
	"kelasd/models/response"
	"kelasd/models/users"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetDetailNewsController(c echo.Context) error {
	newsId, _ := strconv.Atoi(c.Param("newsId"))

	var data = news.News{} // DB

	var response = response.BaseResponse{
		newsId,
		"success",
		data,
	}
	return c.JSON(http.StatusOK, response)
}

func HelloController(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World")
}

func GetNewsController(c echo.Context) error {
	var news []users.User

	result := configs.DB.Preload("News").Find(&news)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			http.StatusInternalServerError,
			result.Error.Error(),
			nil,
		})
	}

	var response = response.BaseResponse{
		http.StatusOK,
		"success",
		news,
	}
	return c.JSON(http.StatusOK, response)
}

func InsertNewsController(c echo.Context) error {
	var news news.News
	c.Bind(&news)

	result := configs.DB.Create(&news)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			http.StatusInternalServerError,
			result.Error.Error(),
			nil,
		})
	}

	var response = response.BaseResponse{
		http.StatusOK,
		"success",
		news,
	}
	return c.JSON(http.StatusOK, response)
}
