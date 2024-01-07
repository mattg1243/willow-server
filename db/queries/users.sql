-- name: CreateUser :one
INSERT INTO users (
    fname, lname, email, salt, "hash", city, nameforheader, phone, "state", street, zip, license, paymentinfo
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1;

-- name: UpdateUser :one
UPDATE users
SET
    fname = $1,
    lname = $2,
    email = $3,
    city = $4,
    nameForHeader = $5,
    phone = $6,
    "state" = $7,
    street = $8,
    zip = $9,
    license = $10,
    paymentInfo = $11,
    updated_at = NOW()
WHERE
    id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;
