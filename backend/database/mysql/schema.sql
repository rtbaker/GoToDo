DROP TABLE IF EXISTS users;

CREATE TABLE users (
    id         INTEGER PRIMARY KEY AUTO_INCREMENT,
    email      VARCHAR(256) UNIQUE,
    password   TEXT NOT NULL,
    created_at TEXT NOT NULL,
    updated_at TEXT NOT NULL
);

DROP TABLE IF EXISTS todo;

CREATE TABLE todo (
    id          INTEGER PRIMARY KEY AUTO_INCREMENT,
    user        INTEGER NOT NULL,
    title       VARCHAR(256) NOT NULL,
    description TEXT NOT NULL,
    priority    INT NOT NULL,
    completed   BOOL NOT NULL DEFAULT 0,
    created_at  TEXT NOT NULL,
    updated_at  TEXT NOT NULL
);

