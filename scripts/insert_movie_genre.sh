#!/usr/bin/env bash
# Usage: ./insert_movie_genre.sh <database.db>
# Assigns genres to movies round-robin: movie 1 -> genre 1, movie 2 -> genre 2, ...
# wrapping back to genre 1 when genres run out. Assumes movies and genres are loaded.
set -euo pipefail

DB="${1:?Usage: $0 <database.db>}"

sqlite3 "$DB" <<'SQL'
BEGIN;
DELETE FROM movies_genres;

WITH
  m AS (SELECT id, ROW_NUMBER() OVER (ORDER BY id) - 1 AS rn FROM movies),
  g AS (SELECT id, ROW_NUMBER() OVER (ORDER BY id) - 1 AS rn FROM genres),
  n AS (SELECT COUNT(*) AS c FROM genres)
INSERT INTO movies_genres (movie_id, genre_id)
SELECT m.id, g.id
FROM m
CROSS JOIN n
JOIN g ON g.rn = m.rn % n.c;
COMMIT;

SELECT COUNT(*) || ' rows in movies_genres' FROM movies_genres;
SQL