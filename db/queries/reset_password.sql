-- name: CreateResetPasswordToken :one
INSERT INTO reset_password (
  user_id,
  reset_token,
  requested_at,
  expires_at
) VALUES (
  $1,
  $2,
  NOW(),
  NOW() + INTERVAL '1 hour'
) RETURNING *;

-- name: GetResetPasswordToken :one
SELECT * FROM reset_password WHERE reset_token = $1 AND expires_at > NOW();

-- name: GetResetPasswordTokenByUser :one
SELECT * FROM reset_password
WHERE user_id = $1
  AND expires_at > NOW();

-- name: DeleteExpiredResetTokens :exec
DELETE FROM reset_password WHERE expires_at <= NOW();