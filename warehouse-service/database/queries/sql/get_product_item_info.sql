select pit.id    as id,
       p.name    as name,
       p.price   as price,
       pit.color as color,
       pit.size  as size
from product_item pit
    join product p on p.id = pit.product_id
where pit.id = any ($1)