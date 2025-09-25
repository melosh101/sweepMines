CREATE TABLE IF NOT EXISTS users
(
    id       serial      not null primary key,
    username text unique not null,
    email    text unique not null,
    password text        not null
)