insert into cart_item (cart_id, product_item_id, quantity, is_selected)
values (:cart_id, :product_item_id, :quantity, :is_selected)
returning id