select id,
       status,
       create_date,
       delivery_date,
       user_id
from orders
where user_id = $1