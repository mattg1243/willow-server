-- name: CreateUser :one
INSERT INTO users (
    id, fname, lname, email, "hash", nameforheader, license,  created_at, updated_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, NOW(), NOW()
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
    updated_at = NOW()
WHERE
    id = $5
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;
