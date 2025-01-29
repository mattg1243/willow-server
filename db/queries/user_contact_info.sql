-- name: CreateUserContactInfo :one
INSERT INTO user_contact_info (
  id, user_id, phone, city, "state", street, zip, paymentInfo, created_at, updated_at
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, NOW(), NOW()
)
RETURNING *;

-- name: GetUserContactInfo :one
SELECT 
  phone,
  city,
  "state",
  street,
  zip,
  paymentInfo::JSON,
  updated_at,
  created_at
FROM user_contact_info 
WHERE user_id = $1;

-- name: UpdateUserContactInfo :one
UPDATE user_contact_info 
SET
  phone = $1,
  city = $2,
  "state" = $3,
  street = $4,
  zip = $5,
  paymentInfo = $6,
  updated_at = NOW()
WHERE
  user_id = $7
RETURNING *;
