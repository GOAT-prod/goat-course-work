select product_name,
       color,
    size,
    count,
    price
from report_item
where factory_id = $1
  and date >= $2
  and date < $3