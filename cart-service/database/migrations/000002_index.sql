-- +goose Up
create index if not exists idx_cart_user_id on cart(user_id);
create index if not exists idx_cart_item_product_item_id on cart_item(product_item_id);

-- +goose Down
drop index if exists idx_cart_user_id;
drop index if exists idx_cart_item_product_item_id;

