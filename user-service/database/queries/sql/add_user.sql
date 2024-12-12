insert into users (username, password, status, role_id, client_id)
values ($1, $2, $3, $4, $5)
returning id