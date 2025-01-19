-- +goose Up
create table if not exists report_item
(
    id           bigserial primary key,
    product_name text not null,
    factory_id   int  not null,
    user_id      int  not null,
    color        text,
    size         int,
    count        int,
    price        decimal
);

create index if not exists idx_report_item_factory_id on report_item (factory_id);
create index if not exists idx_report_item_user_id on report_item (user_id);

-- +goose Down
