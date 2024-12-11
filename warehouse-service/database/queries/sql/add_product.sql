INSERT INTO product (name, brand, price, status, factory_id)
VALUES (:name, :brand, :price, :status, :factory_id)
RETURNING id;