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

func buildPostQuery(includeIDFilter bool) (string, error) {
	var post_filter string
	if includeIDFilter {
		post_filter = "WHERE p.post_id = $1"
	} else {
		post_filter = ""
	}

	// Note that Sprintf and queries are not recommended (sql injection)
	// Note that we are only using this to build our query but leave a variable
	// to insert user input safely
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
	if len(id) != 0 && len(id) != 1 {
		return nil, errors.New("Can only specify at most one ID")
	}

	// ...string lets us pass 0 or 1 IDs
	// Have a pretty complex query to combine post/comment info so we want to use this logic for getting
	// All posts and a single post
	query, err := buildPostQuery(len(id) > 0)
	if err != nil {
		return nil, err
	}

	db := storage.GetDBConnection()
	posts := []Post{}
	if len(id) == 0 {
		err = db.Select(&posts, query)
	} else {
		err = db.Select(&posts, query, id[0])
	}
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

func AddPost(user User, post Post) error {
	db := storage.GetDBConnection()

	tx := db.MustBegin()
	_, err := tx.Exec(`
		INSERT INTO posts (title, url, created_on, creator_id)
		VALUES ($1, $2, NOW(), $3)
	`, post.Title, post.Url, user.UserID)
	tx.Commit()

	return err

}
