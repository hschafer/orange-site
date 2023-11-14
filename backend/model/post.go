package model

import (
	"errors"
	"fmt"
	"orange-site/backend/storage"
)

type Post struct {
	PostID      int    `db:"post_id"`
	Title       string `db:"title"`
	Url         string `db:"url"`
	CreatedOn   string `db:"created_on"`
	CreatorID   int    `db:"creator_id"`
	CreatorName string `db:"username"`
	NumComments int    `db:"count"`
}

func buildPostQuery(id ...string) (string, error) {
	var post_filter string
	if id == nil {
		// Return all (no filter)
		post_filter = ""
	} else if len(id) == 1 {
		// Ensure spaces around where clause
		post_filter = fmt.Sprintf(" WHERE p.post_id = %s ", id[0])
	} else {
		return "", errors.New("Can specify at most one id")
	}

	query := fmt.Sprintf(`
		SELECT
			p.post_id, p.title, p.url, p.created_on, p.creator_id, u.username, COUNT(c.comment_id)
		FROM posts as p
			LEFT JOIN comments AS c ON p.post_id = c.post_id
			LEFT JOIN users as u ON p.creator_id = u.user_id
		%s
		GROUP BY (p.post_id, u.user_id)
		ORDER BY p.created_on DESC;
	`, post_filter)

	return query, nil
}

func getPosts(id ...string) ([]Post, error) {
	// Have a pretty complex query to combine post/comment info so we want to use this logic for getting
	// All posts and a single post
	query, err := buildPostQuery(id...)
	if err != nil {
		return nil, err
	}

	db := storage.GetDBConnection()
	posts := []Post{}
	err = db.Select(&posts, query)
	return posts, err
}

func GetAllPosts() ([]Post, error) {
	posts, err := getPosts()
	return posts, err
}

func GetPost(id string) (*Post, error) {
	posts, err := getPosts(id)
	if err != nil {
		return nil, err
	} else if len(posts) > 1 {
		return nil, errors.New("ID did not uniquely identify post") // Shouldn't happen though
	} else {
		return &posts[0], err
	}
}
