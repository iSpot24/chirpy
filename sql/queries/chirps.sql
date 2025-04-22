-- name: CreateChirp :one
insert into chirps (id, body, user_id, created_at, updated_at)
values ($1, $2, $3, $4, $5)
returning *;

-- name: GetChirps :many
select * from chirps 
where user_id = coalesce(sqlc.narg('user_id'), user_id)
order by created_at asc;

-- name: GetChirpById :one
select * from chirps where id = $1;

-- name: DeleteChirpById :exec
delete from chirps where id = $1;