package repository

import (
	"context"
	"fmt"

	"github.com/purisaurabh/car-rental/internal/repository/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)


type carStore struct{
	BaseRepository
}


type CarStorer interface{
AddCarRepo(ctx context.Context , req model.CarRepo) (int64, error)
}


func NewCarRepo(db *gorm.DB) CarStorer{
	return &carStore{
		BaseRepository : BaseRepository{DB : db},
	}
}


// AddCarRepo is used to add the car into the database
func (cs *carStore) AddCarRepo(ctx context.Context , req model.CarRepo) (int64, error) {
	// insert the car into the database
	result := cs.DB.Create(&req)

	// Check for errors during insertion
	if result.Error != nil {
		zap.S().Errorw("Failed to add car", "error", result.Error)
		return 0, fmt.Errorf("failed to add car: %w", result.Error)
	}

	// Return the inserted car ID
	return req.ID, nil
}
