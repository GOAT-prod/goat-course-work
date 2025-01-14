INSERT INTO product_item (product_id, color, size, weight, quantity)
VALUES (:product_id, :color, :size, :weight, :stock_count)
RETURNING id;