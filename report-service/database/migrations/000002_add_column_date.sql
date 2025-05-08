-- +goose Up
alter table report_item
    add column if not exists date timestamp not null default now();

-- +goose Down