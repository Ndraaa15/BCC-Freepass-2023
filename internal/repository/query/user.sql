-- name: GetUser :one
SELECT
  users.id,
  users.email,
  users.hashed_password,
  users.role_id,
  users.created_at,
  users.updated_at,
  roles.id,
  roles.name
FROM users JOIN roles ON users.role_id = roles.id
WHERE users.id = $1
LIMIT 1;



