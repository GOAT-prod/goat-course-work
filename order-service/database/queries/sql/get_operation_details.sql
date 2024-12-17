select id,
       operation_id,
       type,
       price
from operation_detail
where operation_id = $1