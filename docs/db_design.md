# Database Design

## Movie related

### `movie`

| Column | Notes                                     |
| ------ | ----------------------------------------- |
| id     |                                           |
| name   |                                           |
| year   |                                           |
| url    | Letterboxd URL                            |
| image  | URL to bucket?                            |
| genre  | Many-to-many — resolved via `movie_genre` |

### `genre`

| Column | Notes                        |
| ------ | ---------------------------- |
| id     |                              |
| type   | Genre name (e.g. `action`)   |

### `movie_genre`

| Column   | Notes            |
| -------- | ---------------- |
| movie_id | → `movie.id`     |
| genre_id | → `genre.id`     |

## User related

### `users`

| Column | Notes     |
| ------ | --------- |
| id     |           |
| name   |           |
| pw     | Add later |

**Soon:** friends

## App related

### `user_movie_list`

| Column   | Notes         |
| -------- | ------------- |
| user_id  | → `users.id`  |
| movie_id | → `movie.id`  |

**Consider:** genre

**Soon:**

- Has seen
- Rank
- Multiple lists

## Other
- docs page in bookmark folder
- sqlitebrowser package installed