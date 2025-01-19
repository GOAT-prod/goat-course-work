select o.user_id,
       oi.product_item_id,
       oi.quantity,
       opd.price,
       op.date
from orders o
         join orders_item oi on o.id = oi.orders_id
         join operation op on o.id = op.orders_id
         join operation_detail opd on op.id = opd.operation_id
where op.date >= $1
  and op.date < $2