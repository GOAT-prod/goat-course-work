-- +goose Up
create index if not exists idx_users_client_id on users(client_id);

-- +goose Down
drop index if exists idx_users_client_id;