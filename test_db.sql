-- Assumes DB has already been created

INSERT INTO users (username, password, created_on) VALUES
    ('user1', 'password1', NOW()),
    ('user2', 'password2', NOW()),
    ('user3', 'password3', NOW()),
    ('user4', 'password4', NOW()),
    ('user5', 'password5', NOW()),
    ('user6', 'password6', NOW()),
    ('user7', 'password7', NOW()),
    ('user8', 'password8', NOW()),
    ('user9', 'password9', NOW()),
    ('user10', 'password10',  NOW());


INSERT INTO posts (title, url, created_on, creator_id) VALUES
    ('Post 1 Title', 'http://example1.com/post1', NOW(), 1),
    ('Post 2 Title', 'http://example2.net/post2', NOW(), 3),
    ('Post 3 Title', 'http://example3.org/post3', NOW(), 4),
    ('Post 4 Title', 'http://example4.co/post4', NOW(), 2),
    ('Post 5 Title', 'http://example5.info/post5', NOW(), 4),
    ('Post 6 Title', 'http://example6.tv/post6', NOW(), 3),
    ('Post 7 Title', 'http://example7.biz/post7', NOW(), 6),
    ('Post 8 Title', 'http://example8.me/post8', NOW(), 9),
    ('Post 9 Title', 'http://example9.us/post9', NOW(), 4),
    ('Post 10 Title', 'http://example10.io/post10', NOW(), 10);

INSERT INTO comments (comment, created_on, creator_id, post_id, parent_id) VALUES
    ('Comment 1 on Post 1', NOW(), 1, 1, NULL),
    ('Comment 2 on Post 1', NOW(), 2, 1, NULL),
    ('Reply to Comment 1 on Post 1', NOW(), 3, 1, 1),
    ('Comment 1 on Post 2', NOW(), 1, 2, NULL),
    ('Comment 2 on Post 2', NOW(), 2, 2, NULL),
    ('Comment 1 on Post 3', NOW(), 3, 3, NULL),
    ('Reply to Comment 1 on Post 3', NOW(), 1, 3, 6),
    ('Comment 2 on Post 3', NOW(), 2, 3, NULL),
    ('Comment 1 on Post 4', NOW(), 3, 4, NULL),
    ('Reply to Comment 1 on Post 4', NOW(), 4, 4, 9),
    ('Comment 1 on Post 5', NOW(), 1, 5, NULL),
    ('Comment 2 on Post 5', NOW(), 2, 5, NULL),
    ('Comment 3 on Post 5', NOW(), 3, 5, NULL),
    ('Reply to Comment 2 on Post 5', NOW(), 4, 5, 12),
    ('Comment 1 on Post 6', NOW(), 1, 6, NULL),
    ('Comment 2 on Post 6', NOW(), 2, 6, NULL),
    ('Comment 3 on Post 6', NOW(), 3, 6, NULL),
    ('Reply to Comment 1 on Post 6', NOW(), 4, 6, 15),
    ('Reply to Comment 2 on Post 6', NOW(), 5, 6, 16),
    ('Reply to Comment 1 Reply on Post 1', NOW(), 5, 1, 3),
    ('Reply to Comment 1 Reply on Post 1', NOW(), 5, 1, 3);