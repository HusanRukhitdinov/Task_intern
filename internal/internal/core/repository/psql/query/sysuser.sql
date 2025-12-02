-- name: CreateSysuser :one
INSERT INTO sysusers (
    id,
    status,
    name,
    phone,
    password,
    created_at,
    created_by
    ) VALUES (gen_random_uuid(),$1,$2,$3,$4,$5,$6) RETURNING id;

-- name: GetSysuserByPhone :many
SELECT
    *
    FROM sysusers WHERE status = $1 AND phone=$2;

