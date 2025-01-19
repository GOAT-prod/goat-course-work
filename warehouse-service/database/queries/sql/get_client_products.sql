SELECT
    p.id,
    p.name,
    p.brand,
    p.price,
    p.status,
    p.factory_id
FROM product p
where p.factory_id = any($1)