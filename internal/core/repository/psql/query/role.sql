-- name: UpdateRole :exec
UPDATE roles SET name=$1  WHERE id = $2;


-- name: RoleList :many
SELECT 
    id,
    name,
    created_at
    FROM roles WHERE status=$1;

-- name: CreateRole :exec
INSERT INTO roles(name, status) VALUES ($1,$2);


-- name: CreateSysRoles :exec
INSERT INTO sysuser_roles(id,sysuser_id,role_id) VALUES ($1,$2,$3);