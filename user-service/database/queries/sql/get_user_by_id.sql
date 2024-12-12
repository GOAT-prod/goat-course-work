select id,
       username,
       password,
       status,
       role_id,
       client_id
from users
where id = $1