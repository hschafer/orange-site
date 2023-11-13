package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var schema = `
CREATE TABLE IF NOT EXISTS users (
	user_id serial PRIMARY KEY,
	username VARCHAR ( 50 ) UNIQUE NOT NULL,
	password VARCHAR ( 50 ) NOT NULL,
	email VARCHAR ( 255 ) UNIQUE NOT NULL,
	created_on TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS posts (
    post_id serial PRIMARY KEY,
    title VARCHAR (100) NOT NULL,
    created_on TIMESTAMP NOT NULL,
    creator_id INT,
    FOREIGN KEY (creator_id)
        REFERENCES users (user_id)
    -- TODO Votes
);

CREATE TABLE IF NOT EXISTS comments (
    comment_id serial PRIMARY KEY,
    comment VARCHAR (1000) NOT NULL,
    created_on TIMESTAMP NOT NULL,
    creator_id INT,
    post_id INT,
    FOREIGN KEY (creator_id)
        REFERENCES users (user_id),
    FOREIGN KEY (post_id)
        REFERENCES posts (post_id)
    -- TODO nested comments
    -- TODO Votes
);
`

type User struct {
	UserID    int    `db:"user_id"`
	Username  string `db:"username"`
	Password  string `db:"password"`
	Email     string `db:"email"`
	CreatedOn string `db:"created_on"`
}

type Post struct {
	PostID    int    `db:"post_id"`
	Title     string `db:"title"`
	CreatedOn string `db:"created_on"`
	CreatorID int    `db:"creator_id"`
}

// Handlers
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func posts(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World2!")
}

// Setup
func main() {
	// Set up DB connection
	db, err := sqlx.Connect("postgres", "host=db user=user password=password dbname=db sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	db.MustExec(schema)

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)
	e.GET("/posts", posts)

	// Test DB
	people := []User{}
	db.Select(&people, "SELECT * FROM users;")
	//fmt.Print(people)

	posts := []Post{}
	db.Select(&posts, "SELECT * FROM posts;")
	fmt.Print(posts)

	// Start server
	e.Logger.Fatal(e.Start(":3000"))
}
