UPDATE product_item
SET
    color = :color,
    size = :size,
    weight = :weight,
    quantity = :quantity
WHERE
    id = :id;