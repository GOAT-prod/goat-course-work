-- +goose Up
create table if not exists request
(
    id          serial primary key,
    status      text not null default 0,
    type        text not null default 0,
    update_date date,
    summary     text
);

create table if not exists request_item
(
    id              serial primary key,
    request_id      int,
    product_id      int,
    product_item_id int 
);

-- +goose Down