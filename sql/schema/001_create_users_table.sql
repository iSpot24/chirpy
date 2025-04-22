-- +goose Up
create table if not exists users (
    id UUID primary key,
    email text not null unique,
    created_at timestamp default CURRENT_TIMESTAMP not null,
    updated_at timestamp default CURRENT_TIMESTAMP not null
);

-- +goose Down
drop table users;