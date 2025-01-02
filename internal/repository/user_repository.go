package repository

import (
	"context"
	"fmt"

	"github.com/purisaurabh/car-rental/internal/repository/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type userStore struct {
	BaseRepository
}

type UserStorer interface {
	UserRegistration(ctx context.Context, userReq *model.UserRegistrationRepo) (int64, error)
	UserLogin(ctx context.Context, email string) (model.Users, error)
}

func NewUserRepo(db *gorm.DB) UserStorer{
	return &userStore{
		BaseRepository: BaseRepository{DB: db},
	}
}

func (repo *userStore) UserRegistration(ctx context.Context, user *model.UserRegistrationRepo) (int64, error) {
	// Insert the user into the database using GORM
	result := repo.DB.Create(&user)

	if result.Error != nil {
		zap.S().Errorw("failed to register user", "error", result.Error , "user ID ", user.ID)
		return 0, fmt.Errorf("failed to register user: %w", result.Error)
	}

	// Return the inserted user ID
	return user.ID, nil
}

func (repo *userStore) UserLogin(ctx context.Context, email string) (model.Users, error) {
	var userInfo model.Users

	// Query the user by email using GORM
	err := repo.DB.WithContext(ctx).
		Where("email = ?", email).
		First(&userInfo).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			zap.S().Infow("user not found during login", "email", email)
			return model.Users{}, fmt.Errorf("user not found")
		}
		zap.S().Errorw("failed to execute login query", "email", email, "error", err)
		return model.Users{}, fmt.Errorf("failed to execute login query: %w", err)
	}

	// Return the user information
	return userInfo, nil
}

