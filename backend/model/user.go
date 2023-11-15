package model

import "orange-site/backend/storage"

type User struct {
	UserID    int    `db:"user_id"`
	Username  string `db:"username"`
	Password  string `db:"password"`
	CreatedOn string `db:"created_on"`
}

func FindUser(username string) (*User, error) {
	db := storage.GetDBConnection()

	user := User{}
	err := db.Get(&user, "SELECT * FROM users WHERE username = $1", username)
	return &user, err
}

func AddUser(username string, hashedPassword string) (*User, error) {
	db := storage.GetDBConnection()

	tx := db.MustBegin()
	_, err := tx.Exec(`
		INSERT INTO users (username, password, created_on)
		VALUES ($1, $2, NOW())
	`, username, hashedPassword)
	tx.Commit()

	if err != nil {
		return nil, err
	}

	return FindUser(username)
}
