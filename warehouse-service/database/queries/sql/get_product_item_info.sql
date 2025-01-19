select pit.id       as id,
       p.id         as product_id,
       p.factory_id as factory_id,
       p.name       as name,
       p.price      as price,
       pit.color    as color,
       pit.size     as size,
       pit.weight   as weight,
       pit.quantity as count
from product_item pit
    join product p on p.id = pit.product_id
where pit.id = any ($1)