package model

import "orange-site/backend/storage"

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
