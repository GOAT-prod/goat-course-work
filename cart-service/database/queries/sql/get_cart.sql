select id,
       create_date,
       user_id
from cart
where user_id = $1