update request
set status = $1,
    update_date = now()
where id = $2