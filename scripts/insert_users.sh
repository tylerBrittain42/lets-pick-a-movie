
#!/usr/bin/env bash
# Usage: ./insert_users.sh <database.db> [csv_path]
set -euo pipefail
 
DB="${1:?Usage: $0 <database.db> [csv_path]}"
CSV="${2:-scripts/sample_data/user_list.csv}"
 
sqlite3 "$DB" <<SQL

DROP TABLE IF EXISTS users_staging;
.import --csv $CSV users_staging
 
BEGIN;
DELETE FROM users;
INSERT OR IGNORE INTO users (name)
SELECT Name FROM users_staging;
COMMIT;
 
DROP TABLE users_staging;
 
SELECT COUNT(*) || ' rows in users' FROM users;
SQL
 
