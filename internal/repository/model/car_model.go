package model

type CarRepo struct {
	ID          int64   `gorm:"column:id;primaryKey;autoIncrement"`
	OwnerID     int64   `gorm:"column:owner_id;"`
	Model       string  `gorm:"column:model"`
	RentPerHour float64 `gorm:"column:rent_per_hour"`
	Latitude    string  `gorm:"column:latitude"`
	Longitude   string  `gorm:"column:longitude"`
	IsAvailable bool    `gorm:"column:is_available"`
	CreatedAt   int64   `gorm:"column:created_at"`
	UpdatedAt   int64   `gorm:"column:updated_at"`
}

func (CarRepo) TableName() string {
	return "cars"
}
