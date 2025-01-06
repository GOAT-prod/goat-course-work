select id,
       request_id,
       product_item_id,
       product_id,
       product_item_count
from request_item
where request_id = $1