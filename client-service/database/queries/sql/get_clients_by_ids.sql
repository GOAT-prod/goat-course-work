select id,
       name,
       inn,
       address
from client
where id = any($1)