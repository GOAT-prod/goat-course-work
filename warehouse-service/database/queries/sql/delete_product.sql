UPDATE product
SET
    status = 'deleted'
WHERE
    id = any($1);
