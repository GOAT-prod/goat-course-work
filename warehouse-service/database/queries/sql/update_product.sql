UPDATE product
SET
    name = :name,
    brand = :brand,
    price = :price,
    status = :status,
    factory_id = :factory_id
WHERE
    id = :id;