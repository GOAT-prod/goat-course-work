INSERT INTO image (product_id, url)
VALUES (:product_id, :url)
RETURNING id;
