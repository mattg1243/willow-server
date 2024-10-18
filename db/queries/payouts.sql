-- name: CreatePayout :one
insert into payouts (
  id, user_id, date, amount, client_id
) values (
  $1, $2, $3, $4, $5
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