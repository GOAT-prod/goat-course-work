insert into users (username, password, status, role_id, client_id)
values (:username, :password, :status, :role_id, :client_id)
returning id