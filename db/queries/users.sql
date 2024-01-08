-- name: CreateUser :one
INSERT INTO users (
    fname, lname, email, "hash", city, nameforheader, phone, "state", street, zip, license, paymentinfo, id, created_at, updated_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, NOW(), NOW()
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1;

-- name: UpdateUser :one
UPDATE users
SET
    fname = $1,
    lname = $2,
    city = $3,
    nameForHeader = $4,
    phone = $5,
    "state" = $6,
    street = $7,
    zip = $8,
    license = $9,
    paymentInfo = $10,
    updated_at = NOW()
WHERE
    id = $11
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;
