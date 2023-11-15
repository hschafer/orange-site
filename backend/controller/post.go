package controller

import (
	"fmt"
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

type NewPostInfo struct {
	Title string
	Url   string
}

func NewPost(c echo.Context) error {
	claims, err := ParseToken(c)
	// TODO refresh if logged in?
	if err != nil {
		return c.String(http.StatusUnauthorized, "Invalid token")
	}

	post := new(NewPostInfo)
	err = c.Bind(post)

	if err != nil {
		return c.String(http.StatusBadRequest, "Can not parse request")
	}

	username := fmt.Sprint(claims["username"])
	user, err := model.FindUser(username)

	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid user")
	}

	err = model.AddPost(*user, model.Post{
		Title: post.Title,
		Url:   post.Url,
	})

	if err != nil {
		return c.String(http.StatusInternalServerError, "Unable to make post")
	}

	return c.String(http.StatusOK, "Post submitted!")
}
