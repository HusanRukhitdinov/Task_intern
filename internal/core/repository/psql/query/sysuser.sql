-- name: SysUsers :one
INSERT INTO sysusers (
    id,
    status,
    name,
    phone,
    created_at,
    created_by
    ) VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING id;

-- name: SelectSysUsers :many
SELECT
    *
    FROM sysusers WHERE status = $1 AND phone=$2;
