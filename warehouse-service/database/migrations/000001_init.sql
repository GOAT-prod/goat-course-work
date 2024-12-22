-- +goose Up
create table if not exists product
(
    id          serial primary key,
    name        text    default ''::text not null,
    brand       text    default ''::text not null,
    price       numeric default 0.0      not null,
    material    text    default ''::text not null,
    status      text    default ''::text not null,
    factory_id  integer
);

create table if not exists product_item
(
    id         serial primary key,
    product_id integer references product(id),
    color      text    default ''::text not null,
    size       numeric default 0.0      not null,
    weight     numeric default 0.0      not null,
    quantity   integer default 0        not null
);

create table if not exists material
(
    id   serial primary key,
    name text not null
);

create table if not exists product_material
(
    id          serial  primary key,
    product_id  integer references product(id),
    material_id integer references material(id)
);

create table if not exists image
(
    id         serial primary key,
    product_id integer references product(id),
    url        text default ''::text not null
);

-- +goose Down