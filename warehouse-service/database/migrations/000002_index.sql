-- +goose Up
create index if not exists idx_product_factory_id on product(factory_id);

-- +goose Down
drop index if exists idx_product_factory_id;