SELECT
    pi.id,
    pi.color,
    pi.size,
    pi.weight,
    pi.quantity as stock_count
FROM
    product_item pi
WHERE
    pi.product_id = $1;