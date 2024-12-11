DELETE FROM product_item
WHERE id = any($1);