package model

import "orange-site/backend/storage"

type User struct {
	UserID    int    `db:"user_id"`
	Username  string `db:"username"`
	Password  string `db:"password"`
	Email     string `db:"email"`
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
