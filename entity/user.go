package entity

type User struct {
	Username  string     `db:"username"`
	Password  string     `db:"password"`
	IsActive  bool       `db:"is_active"`
	UserRoles []UserRole `db:"user_roles"`
}
