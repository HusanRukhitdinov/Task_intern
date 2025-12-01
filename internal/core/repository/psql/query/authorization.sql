-- CREATE TABLE users (
--                        id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
--                        name VARCHAR(255) NOT NULL,
--                        email VARCHAR(255) UNIQUE NOT NULL,
--                        password_hash TEXT NOT NULL,
--                        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
--                        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
-- );



-- -- name: RegisterUser :exec
-- INSERT INTO users(
--                   id, email, password_hash, created_at
-- ) VALUES ($1,$2,$3,$4);

-- -- name: SelectOneUserEmail :one
--     SELECT
--         id,
--         password_hash
--         FROM users WHERE email=$1;

-- -- name: LoginOneUser :one
--     SELECT
--         id
--         FROM users WHERE email=$1 AND password_hash=$2;


-- -- name: EditOneUser :exec
-- UPDATE users SET name=$1,email=$2 ,updated_at=$3 WHERE id=$4;

-- -- name: EditOnePassword :exec
-- UPDATE users SET password_hash=$1 WHERE email=$2;

-- -- name: DeleteOneUser :exec
--     DELETE FROM users WHERE id=$1;


-- -- name: SelectOneUser :one
-- SELECT
--     id,
--     name,
--     email,
--     created_at,
--     updated_at
--     FROM users WHERE id=$1;


-- -- name: SelectManyUsers :many
-- SELECT
--     id,
--     name,
--     email,
--     created_at,
--     updated_at
--     FROM users;


