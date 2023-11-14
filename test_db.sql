-- Assumes DB has already been created

INSERT INTO users(username,password,email,created_on) VALUES ('hunter', 'password', 'test@gmail.com', NOW()), ('scott', 'password2', 'test2@gmail.com', NOW());
INSERT INTO posts(title,url,created_on,creator_id) VALUES ('Post1', 'https://www.google.com', NOW(), 1), ('Post2', 'https://www.bing.com', NOW(), 2);
INSERT INTO comments(comment,created_on,creator_id,post_id) VALUES ('this is a comment', NOW(), 2, 1);