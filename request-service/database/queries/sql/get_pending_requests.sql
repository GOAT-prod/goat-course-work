select id,
       status,
       type,
       update_date,
       summary
from request
where status = 'pending'