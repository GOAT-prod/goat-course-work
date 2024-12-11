DELETE FROM product_material
WHERE id = any($1);