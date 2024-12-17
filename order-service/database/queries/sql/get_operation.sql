select id,
    date,
    description,
    orders_id
from operation
where orders_id = $1