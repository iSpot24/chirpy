-- +goose Up
create table if not exists chirps (
    id uuid primary key,
    body text not null,
    user_id uuid not null references users(id) on delete cascade,
    created_at timestamp not null default CURRENT_TIMESTAMP,
    updated_at timestamp not null default CURRENT_TIMESTAMP
);

-- +goose Down
drop table chirps;