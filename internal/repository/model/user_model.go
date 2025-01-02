package model


type UserRegistrationRepo struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement"`
	Name      string    `gorm:"column:name"`
	Email     string    `gorm:"column:email"`
	Password  string    `gorm:"column:password"`
	Mobile    string    `gorm:"column:mobile"`
	Role      string    `gorm:"column:role"`
	CreatedAt int64 	`gorm:"column:created_at"`
	UpdatedAt int64 	`gorm:"column:updated_at"`
}

// override the table name
func (UserRegistrationRepo) TableName() string{
	return "users"
}


type Users struct {
	UserID   int64  `gorm:"column:id"`
	Name     string `gorm:"column:name"`
	Password string `gorm:"column:password"`
	Mobile   string `gorm:"column:mobile"`
	Role     string `gorm:"column:role"`
}

type UserLoginRepo struct {
	UserID   int64  `db:"user_id"`
	Password string `db:"password"`
}
