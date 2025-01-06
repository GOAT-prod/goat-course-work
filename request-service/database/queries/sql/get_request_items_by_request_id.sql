select id,
       request_id,
       product_item_id,
       product_id
from request_item
where request_id = $1