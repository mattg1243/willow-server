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
select 
  payouts.id,
  payouts.user_id,
  payouts.date,
  payouts.amount,
  payouts.created_at,
  payouts.updated_at,
  clients.id as client_id,
  clients.fname as client_fname,
  clients.lname as client_lname
from payouts
left join clients on payouts.client_id = clients.id
where payouts.id = $1;

-- name: GetPayouts :many
select
  payouts.id,
  payouts.user_id,
  payouts.date,
  payouts.amount,
  payouts.created_at,
  payouts.updated_at,
  clients.id as client_id,
  clients.fname as client_fname,
  clients.lname as client_lname
from payouts
left join clients on payouts.client_id = clients.id
where payouts.user_id = $1 or clients.id = $1;

-- name: DeletePayout :exec
delete from payouts
where id = $1;