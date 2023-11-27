-- These queries could and should be split over multiple files
-- corresponding to each individual table in the db

-- Albums
-- name: GetAlbum :one
SELECT * FROM albums
WHERE id = $1 LIMIT 1;

-- name: ListAlbums :many
SELECT * FROM albums
ORDER BY artist;

-- name: CreateAlbum :one
INSERT INTO albums (
  title, artist, price
) values (
  $1, $2, $3
) RETURNING *;

-- name: UpdateAlbum :one
UPDATE albums
SET title = $2, artist = $3, price = $4
WHERE id = $1
RETURNING *;

-- name: DeleteAlbum :exec
DELETE FROM albums
WHERE id = $1;

-- Users
-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUsers :many
SELECT * FROM users
ORDER BY username;

-- name: CreateUser :one
INSERT INTO users (
  username, email
) VALUES (
  $1, $2
) RETURNING *;

-- name: UpdateUser :one
UPDATE users
set username = $2, email = $3
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- Artists
-- name: GetArtist :one
SELECT * FROM artists
WHERE id = $1 LIMIT 1;

-- name: GetArtists :many
SELECT * FROM artists
ORDER BY name;

-- name: CreateArtist :one
INSERT INTO artists (
  name, birthday
) VALUES (
  $1, $2
) RETURNING *;

-- name: UpdateArtist :one
UPDATE artists
set name = $2, birthday = $3
WHERE id =$1
RETURNING *;

-- name: DeleteArtist :exec
DELETE from artists
WHERE id = $1;