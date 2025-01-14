UPDATE product_item
SET
    color = :color,
    size = :size,
    weight = :weight,
    quantity = :stock_count
WHERE
    id = :id;