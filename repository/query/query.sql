-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1;

-- name: CreateUser :one
INSERT INTO users (
                   email, hashedPassword, role
)
VALUES (
        $1, $2, $3
       )
RETURNING *;