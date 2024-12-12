create table if not exists client
(
    id      serial primary key,
    name    text not null default '',
    inn     text,
    address text
)