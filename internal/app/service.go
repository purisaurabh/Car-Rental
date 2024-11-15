package app

import (
	"context"
	"time"

	"github.com/purisaurabh/car-rental/internal/pkg/middleware"
	"github.com/purisaurabh/car-rental/internal/pkg/specs"
	"github.com/purisaurabh/car-rental/internal/repository"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	Repo repository.Repository
}

type Service interface {
	UserRegistration(ctx context.Context, userRequest *specs.UserRegistrationRequest) (int64, error)
	UserLogin(ctx context.Context, userRequest *specs.UserLoginRequest) (specs.UserLoginResponse, error)
}

func NewService(repo repository.Repository) *service {
	return &service{
		Repo: repo,
	}
}

func (userService *service) UserRegistration(ctx context.Context, userRequest *specs.UserRegistrationRequest) (int64, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	userRegistrationRepo := repository.UserRegistrationRepo{
		Name:      userRequest.Name,
		Email:     userRequest.Email,
		Password:  string(hashedPassword),
		Mobile:    userRequest.Mobile,
		Role:      userRequest.Role,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	userID, err := userService.Repo.UserRegistration(ctx, &userRegistrationRepo)
	if err != nil {
		zap.S().Error("unable to register the user : ", err, "for the email : ", userRequest.Email)
	}
	zap.S().Info("user registered successfully with the email : ", userRequest.Email)
	return userID, err
}

func (userService *service) UserLogin(ctx context.Context, userRequest *specs.UserLoginRequest) (specs.UserLoginResponse, error) {
	// check the email in the databae
	userInfo, err := userService.Repo.UserLogin(ctx, userRequest.Email)
	if err != nil {
		zap.S().Error("unable to login the user : ", err, "for the email : ", userRequest.Email)
		return specs.UserLoginResponse{}, err
	}

	// compare the password
	err = bcrypt.CompareHashAndPassword([]byte(userInfo.Password), []byte(userRequest.Password))
	if err != nil {
		zap.S().Error("password does not match for the email : ", userRequest.Email)
		return specs.UserLoginResponse{}, err
	}

	payload := specs.TokenPayload{
		UserID: userInfo.UserID,
		Email:  userRequest.Email,
		Role:   userInfo.Role,
	}

	// create the token
	token, err := middleware.CreateToken(payload)
	if err != nil {
		zap.S().Error("unable to create the token : ", err)
		return specs.UserLoginResponse{}, err
	}

	loginResponse := specs.UserLoginResponse{
		UserID: userInfo.UserID,
		Token:  token,
	}

	zap.S().Info("user logged in successfully with the email : ", userRequest.Email)
	return loginResponse, nil
}
