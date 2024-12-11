SELECT
    id,
    product_id,
    url
FROM image
WHERE product_id = $1;
