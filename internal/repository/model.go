package repository

type UserRegistrationRepo struct {
	Name      string `db:"name"`
	Email     string `db:"email"`
	Password  string `db:"password"`
	Mobile    string `db:"mobile_no"`
	Role      string `db:"role"`
	CreatedAt int64  `db:"created_at"`
	UpdatedAt int64  `db:"updated_at"`
}

type UserInfo struct {
	UserID   int64  `db:"user_id"`
	Name     string `db:"name"`
	Password string `db:"password"`
	Mobile   string `db:"mobile_no"`
	Role     string `db:"role"`
}

type UserLoginRepo struct {
	UserID   int64  `db:"user_id"`
	Password string `db:"password"`
}
