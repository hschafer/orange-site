package model

import "orange-site/backend/storage"

type User struct {
	UserID    int    `db:"user_id"`
	Username  string `db:"username"`
	Password  string `db:"password"`
	CreatedOn string `db:"created_on"`
}

func AuthenticateUser(username string, password string) (bool, error) {
	db := storage.GetDBConnection()

	// TODO need to handle salting passwords

	user := User{}
	err := db.Get(&user, "SELECT * FROM users WHERE username = $1 and password = $2", username, password)
	if err != nil {
		return false, err
	}

	return true, nil
}

func FindUser(username string) (*User, error) {
	db := storage.GetDBConnection()

	user := User{}
	err := db.Get(&user, "SELECT * FROM users WHERE username = $1", username)

	// TODO sanitize inputs
	return &user, err
}

func AddUser(username string, password string) (*User, error) {
	db := storage.GetDBConnection()

	tx := db.MustBegin()
	_, err := tx.Exec(`
		INSERT INTO users (username, password, created_on)
		VALUES ($1, $2, NOW())
	`, username, password)
	tx.Commit()

	if err != nil {
		return nil, err
	}

	return FindUser(username)
}
