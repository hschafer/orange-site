package controller

import (
	"fmt"
	"net/http"
	"orange-site/backend/model"

	"github.com/labstack/echo/v4"
)

func GetComments(c echo.Context) error {
	id := c.Param("id")
	post, err := model.GetComments(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, nil)
	} else {
		return c.JSON(http.StatusOK, post)
	}
}

type CommentData struct {
	Comment string
	PostID  int
}

func PostComment(c echo.Context) error {
	fmt.Println("MADE IT TO COMMENT")
	claims, err := ParseToken(c)
	// TODO refresh if logged in?
	if err != nil {
		return c.String(http.StatusUnauthorized, "Invalid token")
	}

	comment := new(CommentData)
	err = c.Bind(comment)

	if err != nil {
		return c.String(http.StatusBadRequest, "Can not parse request")
	}

	username := fmt.Sprint(claims["username"])

	user, err := model.FindUser(username)

	if err != nil {
		return c.String(http.StatusBadRequest, "Unable to find user")
	}

	// TODO parent comment ID?
	err = model.AddComment(*user, model.Comment{
		Comment: comment.Comment,
		PostID:  comment.PostID,
	})

	if err != nil {
		return c.String(http.StatusInternalServerError, "Unable to post comment")
	}

	fmt.Printf("COMMENT RECEIVED %+v\n", comment)
	return c.String(http.StatusOK, "Comment posted!")

}
