select id,
       cart_id,
       product_item_id,
       quantity,
       is_selected
from cart_item
where cart_id = $1