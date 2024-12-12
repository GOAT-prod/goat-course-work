insert into client (name, inn, address)
VALUES ($1, $2, $3)
returning id