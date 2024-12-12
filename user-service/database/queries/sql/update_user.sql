update users
set username = :username,
    password = :password,
    status   = :status
where id = :id;