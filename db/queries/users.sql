-- name: CreateUser :one
INSERT INTO users (
    id, fname, lname, email, "hash", nameforheader, license, rate, created_at, updated_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, NOW(), NOW()
) RETURNING
    id, 
    fname, 
    lname, 
    email, 
    nameforheader, 
    license, 
    created_at, 
    updated_at;

-- name: GetUser :one
SELECT 
    id, 
    fname, 
    lname, 
    email, 
    nameforheader, 
    license, 
    rate,
    created_at, 
    updated_at 
FROM users
WHERE id = $1;

-- name: GetUserHash :one
SELECT "hash" FROM users
WHERE id = $1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1;

-- name: UpdateUser :one
UPDATE users
SET
    fname = $1,
    lname = $2,
    nameForHeader = $3,
    license = $4,
    rate = $5,
    updated_at = NOW()
WHERE
    id = $6
RETURNING 
    id, 
    fname, 
    lname, 
    email, 
    nameforheader, 
    license, 
    rate,
    created_at, 
    updated_at;

-- name: UpdateUserPassword :exec
UPDATE users
SET
    "hash" = $2
WHERE
    id = $1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;
