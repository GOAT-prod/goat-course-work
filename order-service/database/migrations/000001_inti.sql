-- +goose Up
create table if not exists orders
(
    id            uuid primary key,
    status        text      not null,
    create_date   timestamp not null,
    delivery_date timestamp,
    user_id       int references users (id)
    );

create table if not exists orders_item
(
    id              uuid primary key,
    orders_id       uuid references orders (id),
    product_item_id int references product_item (id),
    quantity        int
    );

create table if not exists operation
(
    id          uuid primary key,
    date        timestamp not null,
    description text,
    orders_id   uuid references orders (id)
    );

create table if not exists operation_detail
(
    id           uuid primary key,
    operation_id uuid references operation (id),
    type         int     not null,
    price        decimal not null
);

-- +goose Down
