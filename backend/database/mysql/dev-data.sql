# Dummy data

DELETE FROM todo;
DELETE FROM users;

# password is "password"

INSERT INTO users
    (email, password, created_at, updated_at)
VALUES 
    ('test1@test.com', '$2a$14$MTcIPu/LwaL7G088oE3xv.GgpaEBFOKlWSWP6ijxXKMsW5Q/dMzbu', NOW(), NOW()),
    ('test2@test.com', '$2a$14$MTcIPu/LwaL7G088oE3xv.GgpaEBFOKlWSWP6ijxXKMsW5Q/dMzbu', NOW(), NOW()),
    ('test3@test.com', '$2a$14$MTcIPu/LwaL7G088oE3xv.GgpaEBFOKlWSWP6ijxXKMsW5Q/dMzbu', NOW(), NOW());

INSERT INTO todo
    (userId, title, description, priority, completed, created_at, updated_at)
VALUES
    ((SELECT id FROM users WHERE email = 'test1@test.com'), 'todo1', 'todo1 description', 1, 0, NOW(), NOW()),
    ((SELECT id FROM users WHERE email = 'test1@test.com'), 'todo2', 'todo1 description', 2, 0, NOW(), NOW()),
    ((SELECT id FROM users WHERE email = 'test1@test.com'), 'todo3', 'todo3 description', 3, 0, NOW(), NOW()),
    ((SELECT id FROM users WHERE email = 'test1@test.com'), 'todo4', 'todo4 description', 1, 1, NOW(), NOW()),
    ((SELECT id FROM users WHERE email = 'test2@test.com'), 'todo5', 'todo5 description', 1, 0, NOW(), NOW()),
    ((SELECT id FROM users WHERE email = 'test2@test.com'), 'todo6', 'todo6 description', 1, 0, NOW(), NOW());
