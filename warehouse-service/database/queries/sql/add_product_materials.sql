INSERT INTO product_material (product_id, material_id)
VALUES (:product_id, :material_id)
RETURNING id;