select id,
       name,
       inn,
       address
from client
where id = $1