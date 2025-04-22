-- name: CreateUser :one
insert into users (id, email, created_at, updated_at, hashed_password)
values (
    $1,
    $2,
    $3,
    $4,
    $5
) returning id, email, is_chirpy_red, created_at, updated_at;

-- name: DeleteUsers :exec
delete from users;

-- name: GetUserByEmail :one
select * from users where email = $1;

-- name: UpdateUser :one
update users
set email = $1, hashed_password = $2, updated_at = Now()
where id = $3
returning id, email, is_chirpy_red, created_at, updated_at;

-- name: MarkUserAsChirpyRed :exec
update users
set is_chirpy_red = true, updated_at = Now()
where id = $1;