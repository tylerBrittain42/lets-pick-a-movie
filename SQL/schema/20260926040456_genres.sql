-- +goose UP
CREATE TABLE genres(
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL
);

-- +goose DOWN
DROP TABLE  genres;