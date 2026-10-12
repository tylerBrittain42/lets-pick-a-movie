
#!/usr/bin/env bash
# Usage: ./insert_genres.sh <database.db> [csv_path]
set -euo pipefail
 
DB="${1:?Usage: $0 <database.db> [csv_path]}"
CSV="${2:-scripts/sample_data/genre_list.csv}"
 
sqlite3 "$DB" <<SQL

DROP TABLE IF EXISTS genres_staging;
.import --csv $CSV genres_staging
 
BEGIN;
DELETE FROM genres;
INSERT OR IGNORE INTO genres (name)
SELECT Name FROM genres_staging;
COMMIT;
 
DROP TABLE genres_staging;
 
SELECT COUNT(*) || ' rows in genres' FROM genres;
SQL
 
