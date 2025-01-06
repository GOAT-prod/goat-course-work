select id,
       status,
       type,
       update_date,
       summary
from request
where id = $1