SELECT
    p.id,
    p.name,
    p.brand,
    p.price,
    p.status,
    p.factory_id
FROM product p
where p.id = any($1)