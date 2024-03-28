-- name: InsertUser :one
INSERT INTO users (
    email,
    hashed_password,
    full_name,
    nim,
    major,
    faculty,
    semester,
    contact,
    role_id
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    $9
) RETURNING id;

-- name: FindUserByEmail :one
SELECT u.id, 
u.email, 
u.hashed_password,
u.full_name,
u.nim,
u.major,
u.faculty,
u.total_sks,
u.semester,
u.contact,
r.name as role_name
FROM users u 
JOIN roles r ON u.role_id = r.id 
WHERE email = $1;






