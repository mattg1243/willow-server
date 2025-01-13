-- name: CreatePayout :one
insert into payouts (
  id, user_id, date, amount, client_id, created_at, updated_at
) values (
  $1, $2, NOW(), $3, $4, NOW(), NOW()
) returning *;

-- name: AddEventToPayout :one
insert into payout_events (
  payout_id, event_id
) values (
  $1, $2
) returning *;

-- name: GetPayout :one
select * 
from payouts
where id = $1;

-- name: GetPayouts :many
select * 
from payouts
where user_id = $1 or client_id = $1;

-- name: DeletePayout :exec
delete from payouts
where id = $1;