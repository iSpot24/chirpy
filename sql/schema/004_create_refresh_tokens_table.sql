-- +goose Up
create table if not exists refresh_tokens (
    token text primary key,
    user_id uuid not null references users(id) on delete cascade,
    expires_at timestamp not null,
    revoked_at timestamp,
    created_at timestamp not null default CURRENT_TIMESTAMP,
    updated_at timestamp not null default CURRENT_TIMESTAMP
);

-- +goose Down
drop table if exists refresh_tokens;