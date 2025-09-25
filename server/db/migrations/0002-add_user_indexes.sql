CREATE INDEX IF NOT EXISTS idx_user_id ON users
    (
     id ASC
        );

CREATE INDEX IF NOT EXISTS idx_user_username ON users
    (
     username ASC
        );

CREATE INDEX IF NOT EXISTS idx_user_email ON users
    (
     email ASC
        );