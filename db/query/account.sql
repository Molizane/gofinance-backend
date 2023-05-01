-- name: CreateAccount :one
INSERT INTO accounts (
  user_id, category_id, title, type, description, value, date
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
) RETURNING *;

-- name: GetAccount :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1;

-- name: GetAccounts :many
SELECT a.id, a.user_id, a.title, a.type, a.description, a.value, a.date, a.created_at, c.title as category_title
FROM accounts a
LEFT JOIN categories c ON c.id = a.category_id
WHERE a.user_id = @user_id
AND a.type = @type
AND LOWER(a.title) LIKE CONCAT('%', LOWER(@title::text), '%')
AND LOWER(a.description) LIKE CONCAT('%', LOWER(@description::text), '%')
AND a.category_id = @category_id
AND a.date = @date;

-- name: GetAccountsByUserIdAndType :many
SELECT a.id, a.user_id, a.title, a.type, a.description, a.value, a.date, a.created_at, c.title as category_title
FROM accounts a
LEFT JOIN categories c ON c.id = a.category_id
WHERE a.user_id = @user_id
AND a.type = @type;

-- name: GetAccountsByUserIdAndTypeAndCategoryId :many
SELECT a.id, a.user_id, a.title, a.type, a.description, a.value, a.date, a.created_at, c.title as category_title
FROM accounts a
LEFT JOIN categories c ON c.id = a.category_id
WHERE a.user_id = @user_id
AND a.type = @type
AND a.category_id = @category_id;

-- name: GetAccountsByUserIdAndTypeAndCategoryIdAndTitle :many
SELECT a.id, a.user_id, a.title, a.type, a.description, a.value, a.date, a.created_at, c.title as category_title
FROM accounts a
LEFT JOIN categories c ON c.id = a.category_id
WHERE a.user_id = @user_id
AND a.type = @type
AND a.category_id = @category_id
AND LOWER(a.title) LIKE CONCAT('%', LOWER(@title::text), '%');

-- name: GetAccountsByUserIdAndTypeAndCategoryIdAndDescription :many
SELECT a.id, a.user_id, a.title, a.type, a.description, a.value, a.date, a.created_at, c.title as category_title
FROM accounts a
LEFT JOIN categories c ON c.id = a.category_id
WHERE a.user_id = @user_id
AND a.type = @type
AND a.category_id = @category_id
AND LOWER(a.description) LIKE CONCAT('%', LOWER(@description::text), '%');

-- name: GetAccountsByUserIdAndTypeAndDescription :many
SELECT a.id, a.user_id, a.title, a.type, a.description, a.value, a.date, a.created_at, c.title as category_title
FROM accounts a
LEFT JOIN categories c ON c.id = a.category_id
WHERE a.user_id = @user_id
AND a.type = @type
AND LOWER(a.description) LIKE CONCAT('%', LOWER(@description::text), '%');

-- name: GetAccountsByUserIdAndTypeAndTitle :many
SELECT a.id, a.user_id, a.title, a.type, a.description, a.value, a.date, a.created_at, c.title as category_title
FROM accounts a
LEFT JOIN categories c ON c.id = a.category_id
WHERE a.user_id = @user_id
AND a.type = @type
AND LOWER(a.title) LIKE CONCAT('%', LOWER(@title::text), '%');

-- name: GetAccountsByUserIdAndTypeAndCategoryIdAndTitleAndDescription :many
SELECT a.id, a.user_id, a.title, a.type, a.description, a.value, a.date, a.created_at, c.title as category_title
FROM accounts a
LEFT JOIN categories c ON c.id = a.category_id
WHERE a.user_id = @user_id
AND a.type = @type
AND a.category_id = @category_id
AND LOWER(a.title) LIKE CONCAT('%', LOWER(@title::text), '%')
AND LOWER(a.description) LIKE CONCAT('%', LOWER(@description::text), '%');

-- name: GetAccountsByUserIdAndTypeAndDate :many
SELECT a.id, a.user_id, a.title, a.type, a.description, a.value, a.date, a.created_at, c.title as category_title
FROM accounts a
LEFT JOIN categories c ON c.id = a.category_id
WHERE a.user_id = @user_id
AND a.type = @type
AND a.date = @date;

-- name: GetAccountsReports :one
SELECT COALLESCE(SUM(value), 0) AS sum_value
FROM accounts
WHERE user_id = $1
AND type = $2;

-- name: GetAccountsGraph :one
SELECT COUNT(*)
FROM accounts
WHERE user_id = $1 and type = $2;

-- name: UpdateAccount :one
UPDATE accounts
SET title = $2, description = $3, value = $4
WHERE id = $1
RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM accounts
WHERE id = $1;
