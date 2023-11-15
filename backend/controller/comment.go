package controller

import (
	"database/sql"
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
	Comment  string
	PostID   int
	ParentID *int
}

func PostComment(c echo.Context) error {
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

	parentID := sql.NullInt64{}
	if comment.ParentID != nil {
		parentID.Int64 = int64(*comment.ParentID)
		parentID.Valid = true
	}
	err = model.AddComment(*user, model.Comment{
		Comment:  comment.Comment,
		PostID:   comment.PostID,
		ParentID: parentID,
	})

	if err != nil {
		return c.String(http.StatusInternalServerError, "Unable to post comment")
	}

	return c.String(http.StatusOK, "Comment posted!")

}
