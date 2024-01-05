-- name: CreateUser :one
INSERT INTO users (
    username, fname, lname, email, salt, hash, city, nameforheader, phone, state, street, zip, license, paymentinfo
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1;

-- name: UpdateUser :exec
UPDATE users
SET
    username = $2,
    fname = $3,
    lname = $4,
    email = $5,
    salt = $6,
    hash = $7,
    city = $8,
    nameforheader = $9,
    phone = $10,
    state = $11,
    street = $12,
    zip = $13,
    license = $14,
    paymentinfo = $15,
    updated_at = NOW()
WHERE
    id = $1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;
