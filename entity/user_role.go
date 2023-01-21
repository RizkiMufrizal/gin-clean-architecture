package entity

import "github.com/google/uuid"

type UserRole struct {
	Id       uuid.UUID `db:"user_role_id"`
	Role     string    `db:"role"`
	Username string    `db:"username"`
}
