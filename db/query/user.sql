-- name: CreateUser :one
INSERT INTO users (
  username, password, email
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: UpdateUser :one
UPDATE users
SET email = $1
WHERE UPPER(username) = UPPER(@username::text)
RETURNING *;

-- name: UpdatePassword :one
UPDATE users
SET password = $1
WHERE UPPER(username) = UPPER(@username::text)
RETURNING *;

-- name: GetUser :one
SELECT *
FROM users
WHERE UPPER(username) = UPPER(@username::text) LIMIT 1;

-- name: GetUserById :one
SELECT *
FROM users
WHERE id = $1 LIMIT 1;
