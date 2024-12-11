SELECT
    pm.id,
    pm.product_id,
    m.name
FROM product_material pm
JOIN material m ON pm.material_id = m.id
WHERE pm.product_id = $1;
