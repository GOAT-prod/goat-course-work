INSERT INTO product_item (product_id, color, size, weight, quantity)
VALUES (:product_id, :color, :size, :weight, :quantity)
RETURNING id;