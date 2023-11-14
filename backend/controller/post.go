package controller

import (
	"log"
	"net/http"
	"orange-site/backend/model"

	"github.com/labstack/echo/v4"
)

func GetAllPosts(c echo.Context) error {
	posts, err := model.GetAllPosts()
	if err != nil {
		log.Fatalln(err)
	}

	return c.JSON(http.StatusOK, posts)
}

func GetPost(c echo.Context) error {
	id := c.Param("id")
	post, err := model.GetPost(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, nil)
	} else {
		return c.JSON(http.StatusOK, post)
	}
}
