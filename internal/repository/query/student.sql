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





