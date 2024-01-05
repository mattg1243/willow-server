-- These queries could and should be split over multiple files
-- corresponding to each individual table in the db

-- Users
-- name: GetUser :one
SELECT * FROM users
WHERE id = $1;

-- Clients
-- name: CreateClient :one
-- INSERT INTO clients (
--   "user", fname, lname, email, rate 
-- ) VALUES (
--   $1, $2, $3, $4, $5
-- ) RETURNING *;

-- name: GetClient :one

-- name: GetClients :many

-- name: UpdateClient :one

-- name: DeleteClient :exec


-- Events
-- name: CreateEvent :one

-- name: GetEvent :one

-- name: GetEvents :many

-- name: UpdateEvent :one

-- name: DeleteEvent :exec

-- Albums
-- -- name: GetAlbum :one
-- SELECT * FROM albums
-- WHERE id = $1 LIMIT 1;

-- -- name: ListAlbums :many
-- SELECT * FROM albums
-- ORDER BY artist;

-- -- name: CreateAlbum :one
-- INSERT INTO albums (
--   title, artist, price
-- ) values (
--   $1, $2, $3
-- ) RETURNING *;

-- -- name: UpdateAlbum :one
-- UPDATE albums
-- SET title = $2, artist = $3, price = $4
-- WHERE id = $1
-- RETURNING *;

-- -- name: DeleteAlbum :exec
-- DELETE FROM albums
-- WHERE id = $1;

-- -- Users
-- -- name: GetUser :one
-- SELECT id, username, email, balance FROM users
-- WHERE id = $1 LIMIT 1;

-- -- name: GetUserByUsername :one
-- SELECT id, username, email, balance FROM users
-- WHERE username = $1 LIMIT 1;

-- -- name: GetUserWithHash :one
-- SELECT * FROM users
-- WHERE username = $1 LIMIT 1;

-- -- name: GetUsers :many
-- SELECT id, username, email, balance FROM users
-- ORDER BY username;

-- -- name: CreateUser :one
-- INSERT INTO users (
--   username, "hash", email, balance
-- ) VALUES (
--   $1, $2, $3, $4
-- ) RETURNING *;

-- -- name: UpdateUser :one
-- UPDATE users
-- set username = $2, email = $3, balance = $4
-- WHERE id = $1
-- RETURNING *;

-- -- name: DeleteUser :exec
-- DELETE FROM users
-- WHERE id = $1;

-- -- Artists
-- -- name: GetArtist :one
-- SELECT * FROM artists
-- WHERE id = $1 LIMIT 1;

-- -- name: GetArtists :many
-- SELECT * FROM artists
-- ORDER BY name;

-- -- name: CreateArtist :one
-- INSERT INTO artists (
--   name, birthday
-- ) VALUES (
--   $1, $2
-- ) RETURNING *;

-- -- name: UpdateArtist :one
-- UPDATE artists
-- set name = $2, birthday = $3
-- WHERE id =$1
-- RETURNING *;

-- -- name: DeleteArtist :exec
-- DELETE from artists
-- WHERE id = $1;

-- -- Purchases
-- -- name: CreatePurchase :one
-- INSERT INTO purchases (
--   "user", album, "date"
-- ) VALUES (
--   $1, $2, $3
-- ) RETURNING *;

-- -- name: GetPurchases :many
-- SELECT * FROM purchases
-- ORDER BY id;