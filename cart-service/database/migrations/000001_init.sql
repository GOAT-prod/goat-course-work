-- +goose Up
create table if not exists cart
(
    id          serial primary key,
    create_date timestamp,
    user_id     int
    );

create table if not exists cart_item
(
    id              serial primary key,
    cart_id         int references cart (id),
    product_item_id int,
    quantity        int,
    is_selected     bool not null default false
);

-- +goose Down