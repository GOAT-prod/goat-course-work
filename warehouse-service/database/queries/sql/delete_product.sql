UPDATE product
SET
    is_deleted = TRUE
WHERE
    id = any($1);
