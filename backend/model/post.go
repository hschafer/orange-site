package model

import (
	"orange-site/backend/storage"
)

type Post struct {
	PostID    int    `db:"post_id"`
	Title     string `db:"title"`
	CreatedOn string `db:"created_on"`
	CreatorID int    `db:"creator_id"`
}

func GetAllPosts() ([]Post, error) {
	db := storage.GetDBConnection()

	posts := []Post{}
	db.Select(&posts, "SELECT * FROM posts;")
	return posts, nil
}

func GetPost(id string) (*Post, error) {
	db := storage.GetDBConnection()

	var post Post
	err := db.Get(&post, "SELECT * FROM posts WHERE post_id = $1", id)

	if err != nil {
		return nil, err
	} else {
		return &post, err
	}
}
