-- +goose UP
CREATE TABLE movies (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    year INTEGER NOT NULL,
    url TEXT NOT NULL UNIQUE,
    image TEXT DEFAULT "none"
);

-- +goose DOWN
DROP TABLE movies;