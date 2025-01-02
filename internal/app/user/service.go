package user

import (
	"context"
	"fmt"
	"time"

	"github.com/purisaurabh/car-rental/internal/pkg/middleware"
	"github.com/purisaurabh/car-rental/internal/pkg/specs"
	"github.com/purisaurabh/car-rental/internal/repository"
	"github.com/purisaurabh/car-rental/internal/repository/model"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	userRepo repository.UserStorer
}

type Service interface {
	UserRegistration(ctx context.Context, userRequest *specs.UserRegistrationRequest) (int64, error)
	UserLogin(ctx context.Context, userRequest *specs.UserLoginRequest) (specs.UserLoginResponse, error)
}

func NewService(userRepo repository.UserStorer) Service {
	return &service{
		userRepo: userRepo,
	}
}

func (userService *service) UserRegistration(ctx context.Context, userRequest *specs.UserRegistrationRequest) (int64, error) {
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		zap.S().Errorw("Failed to hash password", "email", userRequest.Email, "error", err)
		return 0, fmt.Errorf("failed to hash password: %w", err)
	}

	// Create the user repository model from the request
	user := model.UserRegistrationRepo{
		Name:     userRequest.Name,
		Email:    userRequest.Email,
		Password: string(hashedPassword),
		Mobile:   userRequest.Mobile,
		Role:     userRequest.Role,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	// Call the repository method to insert the user
	userID, err := userService.userRepo.UserRegistration(ctx, &user)
	if err != nil {
		zap.S().Errorw("Failed to register user", "email", userRequest.Email, "error", err)
		return 0, fmt.Errorf("failed to register user: %w", err)
	}

	// Log the successful registration
	zap.S().Infow("User registered successfully", "email", userRequest.Email, "userID", userID)

	// Return the user ID
	return userID, nil
}


func (userService *service) UserLogin(ctx context.Context, userRequest *specs.UserLoginRequest) (specs.UserLoginResponse, error) {
	// Retrieve user data from the repository
	userInfo, err := userService.userRepo.UserLogin(ctx, userRequest.Email)
	if err != nil {
		zap.S().Errorw("User login failed", "email", userRequest.Email, "error", err)
		return specs.UserLoginResponse{}, fmt.Errorf("user login failed: %w", err)
	}

	// Compare provided password with the hashed password
	if err = bcrypt.CompareHashAndPassword([]byte(userInfo.Password), []byte(userRequest.Password)); err != nil {
		zap.S().Warnw("Password mismatch", "email", userRequest.Email)
		return specs.UserLoginResponse{}, fmt.Errorf("invalid credentials")
	}

	// Prepare payload for token generation
	payload := specs.TokenPayload{
		UserID: userInfo.UserID,
		Email:  userRequest.Email,
		Role:   userInfo.Role,
	}

	// Generate token
	token, err := middleware.CreateToken(payload)
	if err != nil {
		zap.S().Errorw("Failed to generate token", "email", userRequest.Email, "error", err)
		return specs.UserLoginResponse{}, fmt.Errorf("failed to generate token: %w", err)
	}

	response := specs.UserLoginResponse{
		UserID: userInfo.UserID,
		Token:  token,
	}

	zap.S().Infow("User logged in successfully", "email", userRequest.Email)
	return response, nil
}
