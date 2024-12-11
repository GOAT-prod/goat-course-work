DELETE FROM image
WHERE id = any($1);
