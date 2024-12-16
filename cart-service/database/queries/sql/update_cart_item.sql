update cart_item
set quantity    = :quantity,
    is_selected = :is_selected
where id = :id