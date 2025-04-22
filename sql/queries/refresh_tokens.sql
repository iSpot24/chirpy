-- name: CreateRefreshToken :exec
insert into refresh_tokens (token, user_id, expires_at)
values ($1, $2, $3);

-- name: GetRefreshToken :one
select * from refresh_tokens where token = $1;

-- name: UpdateRevokedAt :exec
update refresh_tokens
set revoked_at = $2, updated_at = $3
where token = $1;