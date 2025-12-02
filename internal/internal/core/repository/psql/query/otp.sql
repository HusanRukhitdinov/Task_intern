-- name: InsertOTP :one
INSERT INTO otp (
    id,
    email,
    code,
    status,
    expires_at
) VALUES (gen_random_uuid(), $1, $2, $3, $4) 
RETURNING id;

-- name: GetOTPByID :one
SELECT id, email, code, status, expires_at
FROM otp
WHERE id = sqlc.arg(id)::uuid AND status = sqlc.arg(status);

-- name: UpdateOTPStatus :exec
UPDATE otp SET status = sqlc.arg(status) WHERE id = sqlc.arg(id)::uuid;
