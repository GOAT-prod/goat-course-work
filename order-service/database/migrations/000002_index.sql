-- +goose Up
create index if not exists idx_orders_user_id on orders(user_id);
create index if not exists idx_orders_item_product_item_id on orders_item(product_item_id);

-- +goose Down
drop index if exists idx_orders_user_id;
drop index if exists idx_orders_item_product_item_id;