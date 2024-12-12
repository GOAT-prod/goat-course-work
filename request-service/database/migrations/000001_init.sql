create table if not exists request
(
    id          serial primary key,
    status      int not null default 0,
    type        int not null default 0,
    update_date date,
    summary     text
);

create table if not exists request_item
(
    id              serial primary key,
    request_id      int references request (id),
    product_id      int references product (id),
    product_item_id int references product_item (id)
);