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

-- name: UpdateAlbum :exec
UPDATE albums
SET title = $2, artist = $3, price = $4
WHERE id = $1;

-- name: DeleteAlbum :exec
DELETE FROM albums
WHERE id = $1;