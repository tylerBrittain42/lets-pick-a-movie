-- +goose Up
CREATE TABLE movies_genres (
    movie_id,
    genre_id,
    PRIMARY KEY(movie_id, genre_id),
    FOREIGN KEY(movie_id) references movies(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY(genre_id) references genres(id) ON DELETE CASCADE ON UPDATE CASCADE
);

-- +goose Down
DROP TABLE movies_genres;