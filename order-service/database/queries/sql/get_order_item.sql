select id,
       orders_id,
       product_item_id,
       quantity
from orders_item
where orders_id = $1