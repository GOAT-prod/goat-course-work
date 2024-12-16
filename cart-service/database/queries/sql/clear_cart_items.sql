delete
from cart_item
where cart_id = $1
  and is_selected = true;