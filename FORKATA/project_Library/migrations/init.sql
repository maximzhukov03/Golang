CREATE TABLE IF NOT EXISTS users (
    id      VARCHAR PRIMARY KEY,
    name    VARCHAR NOT NULL,
    email   VARCHAR
);

CREATE TABLE IF NOT EXISTS authors (
    id         VARCHAR PRIMARY KEY,
    name       VARCHAR NOT NULL,
    popularity INTEGER DEFAULT 0
);

CREATE TABLE IF NOT EXISTS books (
    id        VARCHAR PRIMARY KEY,
    title     VARCHAR NOT NULL,
    author_id VARCHAR REFERENCES authors(id) ON DELETE CASCADE,
    user_id   VARCHAR REFERENCES users(id) ON DELETE SET NULL
);