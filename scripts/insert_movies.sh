
#!/usr/bin/env bash
# Usage: ./import_movies.sh <database.db> [csv_path]
set -euo pipefail
 
DB="${1:?Usage: $0 <database.db> [csv_path]}"
CSV="${2:-scripts/sample_data/movie_list.csv}"
 
sqlite3 "$DB" <<SQL

DROP TABLE IF EXISTS movies_staging;
.import --csv $CSV movies_staging
 
BEGIN;
DELETE FROM movies;
INSERT OR IGNORE INTO movies (name, year, url)
SELECT Name, Year, url FROM movies_staging;
COMMIT;
 
DROP TABLE movies_staging;
 
SELECT COUNT(*) || ' rows in movies' FROM movies;
SQL
 
