-- name: InsertUser :exec
INSERT INTO users (
    email,
    name,
    password,
    status
) VALUES ($1, $2, $3, $4);

-- name: GetUserByEmail :one
SELECT id, email, name, password, status, created_at
FROM users
WHERE email = $1 AND status = $2;
