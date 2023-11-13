-- This file is not programmatically run, just a log of the table creations

CREATE TABLE users (
	user_id serial PRIMARY KEY,
	username VARCHAR ( 50 ) UNIQUE NOT NULL,
	password VARCHAR ( 50 ) NOT NULL,
	email VARCHAR ( 255 ) UNIQUE NOT NULL,
	created_on TIMESTAMP NOT NULL
);

CREATE TABLE posts (
    post_id serial PRIMARY KEY,
    title VARCHAR (100) NOT NULL,
    created_on TIMESTAMP NOT NULL,
    creator_id INT,
    FOREIGN KEY (creator_id)
        REFERENCES users (user_id)
    -- TODO Votes
);

CREATE TABLE comments (
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