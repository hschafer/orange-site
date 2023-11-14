package model

import (
	"database/sql"
	"orange-site/backend/storage"
)

type Comment struct {
	CommentID int           `db:"comment_id"`
	Comment   string        `db:"comment"`
	CreatedOn string        `db:"created_on"`
	CreatorID int           `db:"creator_id"`
	Username  string        `db:"username"`
	PostID    int           `db:"post_id"`
	ParentID  sql.NullInt64 `db:"parent_id"`
	Children  []*Comment
}

func GetComments(id string) ([]*Comment, error) {
	db := storage.GetDBConnection()

	comments := []*Comment{}
	err := db.Select(&comments, `
		SELECT c.comment_id, c.comment, c.created_on, c.creator_id, c.post_id, c.parent_id, u.username
		FROM comments AS c
		LEFT JOIN users AS u ON c.creator_id = u.user_id
		WHERE c.post_id = $1
		ORDER BY c.created_on DESC
	`, id)

	if err != nil {
		return nil, err
	}

	// Currently all comments are flat, need to add hierarchy of parents
	// First build an index of CommentID -> Comment
	index := make(map[int]*Comment)
	for _, comment := range comments {
		index[comment.CommentID] = comment
	}

	// Build up result adding comment to the top-level results or to their parent
	result := []*Comment{}
	for _, comment := range comments {
		if comment.ParentID.Valid {
			parentID := comment.ParentID.Int64
			parent := index[int(parentID)]
			parent.Children = append(parent.Children, comment)
		} else { // nil
			result = append(result, comment)
		}
	}
	return result, err
}
