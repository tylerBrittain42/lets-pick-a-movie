-- +goose Up
CREATE TABLE users_movies (
    user_id,
    movie_id,
    PRIMARY KEY(user_id, movie_id),
    FOREIGN KEY(user_id) references users(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY(movie_id) references movies(id) ON DELETE CASCADE ON UPDATE CASCADE
);

-- +goose Down
DROP TABLE users_movies;
