-- +goose Up
create table if not exists user_role
(
    id   serial primary key,
    name text not null default ''
);

create index if not exists idx_user_role_name on user_role (name);

insert into user_role (name)
values ('admin'),
       ('shop'),
       ('factory');

create table if not exists users
(
    id        serial primary key,
    username  text not null default '',
    password  text not null default '',
    status    int,
    role_id   int references user_role (id),
    client_id int references client (id)
);

create index if not exists idx_users_username on users (username);

-- +goose Down