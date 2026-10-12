#!/usr/bin/env bash
# Usage: ./insert_movie_user.sh <database.db>
# Assigns users to movies round-robin: movie 1 -> user 1, movie 2 -> user 2, ...
# wrapping back to user 1 when users run out. Assumes movies and users are loaded.
set -euo pipefail

DB="${1:?Usage: $0 <database.db>}"

sqlite3 "$DB" <<'SQL'
BEGIN;
DELETE FROM users_movies;

WITH
  m AS (SELECT id, ROW_NUMBER() OVER (ORDER BY id) - 1 AS rn FROM movies),
  g AS (SELECT id, ROW_NUMBER() OVER (ORDER BY id) - 1 AS rn FROM users),
  n AS (SELECT COUNT(*) AS c FROM users)
INSERT INTO users_movies (movie_id, user_id)
SELECT m.id, g.id
FROM m
CROSS JOIN n
JOIN g ON g.rn = m.rn % n.c;
COMMIT;

SELECT COUNT(*) || ' rows in users_movies' FROM users_movies;
SQL