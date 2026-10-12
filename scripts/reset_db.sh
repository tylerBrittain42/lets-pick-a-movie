#!/usr/bin/env bash

db_name="database.db"
should_reset=false
if [ "${1:-}" = "--hard" ]; then
    should_reset=true
fi

set -euo pipefail
 
# recreate database entirely
if [ "$should_reset" = true ]; then
    if [ -e "$db_name" ]; then
        rm "$db_name"
        echo "$db_name" deleted
    fi

    touch "$db_name"
    echo file "$db_name" created
    cd SQL/schema
    goose sqlite3 ../../database.db up
    cd ..
    cd ..
fi

# run scripts without relations here
./scripts/insert_genres.sh "$db_name"
./scripts/insert_movies.sh "$db_name"
./scripts/insert_users.sh "$db_name"

# scripts with relations
./scripts/insert_movie_genre.sh "$db_name"
./scripts/insert_movie_user.sh "$db_name"

echo sample data loading complete
