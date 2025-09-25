-- name: GetUserById :one
select *
from users
WHERE id = $1
LIMIT 1;

-- name: GetUserByEmail :one
SELECT *
from users
WHERE email = $1
limit 1;

-- name: GetUserByName :one
SELECT *
FROM users
WHERE username = $1
limit 1;


-- name: CreateUser :one
INSERT INTO users (username, email, password)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateUser :exec
UPDATE users
set username = $2,
    email    = $3
where id = $1;

-- name: ChangeUserPass :exec
UPDATE users
set password = $2
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users
where id = $1;



