package storage

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
)

var schema = `
CREATE TABLE IF NOT EXISTS users (
	user_id serial PRIMARY KEY,
	username VARCHAR ( 50 ) UNIQUE NOT NULL,
	password VARCHAR ( 100 ) NOT NULL,
	created_on TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS posts (
    post_id serial PRIMARY KEY,
    title VARCHAR (100) NOT NULL,
	url VARCHAR (100) NOT NULL,
    created_on TIMESTAMP NOT NULL,
    creator_id INT,
    FOREIGN KEY (creator_id)
        REFERENCES users (user_id)
    -- TODO Votes
);

CREATE TABLE IF NOT EXISTS comments (
    comment_id serial PRIMARY KEY,
    comment TEXT NOT NULL,
    created_on TIMESTAMP NOT NULL,
    creator_id INT,
    post_id INT,
	parent_id INT,
    FOREIGN KEY (creator_id)
        REFERENCES users (user_id),
    FOREIGN KEY (post_id)
        REFERENCES posts (post_id),
	FOREIGN KEY (parent_id)
		REFERENCES comments (comment_id)
    -- TODO nested comments
    -- TODO Votes
);
`

type User struct {
	UserID    int    `db:"user_id"`
	Username  string `db:"username"`
	Password  string `db:"password"`
	CreatedOn string `db:"created_on"`
}

var DB *sqlx.DB

func InitDBConnection() *sqlx.DB {
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")

	connectionInfo := fmt.Sprintf("host=db user=%s password=%s dbname=%s sslmode=disable",
		dbUser, dbPassword, dbName)

	// Set up DB connection
	db, err := sqlx.Connect("postgres", connectionInfo)
	if err != nil {
		log.Fatalln(err)
	}

	db.MustExec(schema)
	DB = db
	return DB
}

func GetDBConnection() *sqlx.DB {
	return DB
}
