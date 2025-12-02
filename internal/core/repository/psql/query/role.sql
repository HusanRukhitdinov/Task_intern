-- name: UpdateRole :exec
UPDATE roles SET name=$1  WHERE id = $2;


-- name: RoleList :many
SELECT 
    id,
    name,
    created_at
    FROM roles WHERE status=$1;

-- name: CreateRole :exec
INSERT INTO roles(id, name, status, created_at) VALUES (gen_random_uuid(), $1, $2, CURRENT_TIMESTAMP);

-- name: GetRoleById :one
SELECT id, name, status FROM roles WHERE id = $1 AND status = $2;

-- name: CreateSysRoles :exec
INSERT INTO sysuser_roles(id,sysuser_id,role_id) VALUES (gen_random_uuid(),$1,$2);