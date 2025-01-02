package app

import (
	"github.com/purisaurabh/car-rental/internal/app/car"
	"github.com/purisaurabh/car-rental/internal/app/user"
	"github.com/purisaurabh/car-rental/internal/repository"
	"gorm.io/gorm"
)

type Dependencies struct {
	UserService user.Service
	CarService car.Service 	
}


func NewServices(db *gorm.DB) Dependencies{
	// initialize the repo services
	userRepo := repository.NewUserRepo(db)
	carRepo := repository.NewCarRepo(db)

	userService := user.NewService(userRepo)
	carService := car.NewService(carRepo)

	return Dependencies{
		UserService: userService,
		CarService: carService,
	}
}