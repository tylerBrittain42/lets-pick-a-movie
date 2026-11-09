-- name: GetMovie :one
SELECT movies.name,
    movies.year,
    movies.url,
    json_group_array(genres.name)
FROM movies
    INNER JOIN movies_genres as mg ON movies.id = mg.movie_id
    INNER JOIN genres on genres.id = mg.genre_id
WHERE movies.name = ?
GROUP BY movies.id