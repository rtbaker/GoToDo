DROP TABLE IF EXISTS users;

CREATE TABLE users (
    id         INTEGER PRIMARY KEY AUTO_INCREMENT,
    email      VARCHAR(256) UNIQUE,
    password   TEXT NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL
);

DROP TABLE IF EXISTS todo;

CREATE TABLE todo (
    id          INTEGER PRIMARY KEY AUTO_INCREMENT,
    userId      INTEGER NOT NULL,
    title       VARCHAR(256) NOT NULL,
    description TEXT NOT NULL,
    priority    INT NOT NULL,
    completed   BOOL NOT NULL DEFAULT 0,
    created_at  DATETIME NOT NULL,
    updated_at  DATETIME NOT NULL
);

