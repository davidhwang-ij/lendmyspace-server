package repository

const (
	GetUserQuery    = `SELECT id, user_name, first_name, last_name, email FROM users WHERE id = $1 LIMIT 1`
	ListUsersQuery  = `SELECT user_name, first_name, last_name, email FROM users ORDER BY id LIMIT $1 OFFSET $2`
	CreateUserQuery = `INSERT INTO users(user_name, first_name, last_name, email, password) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	UpdateUserQuery = `UPDATE users SET user_name = $2 WHERE id = $1 RETURNING id`
	DeleteUserQuery = `DELETE FROM users WHERE id = $1`
)
