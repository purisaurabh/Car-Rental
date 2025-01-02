package car

import (
	"context"
	"time"

	"github.com/purisaurabh/car-rental/internal/pkg/specs"
	"github.com/purisaurabh/car-rental/internal/repository"
	"github.com/purisaurabh/car-rental/internal/repository/model"
	"go.uber.org/zap"
)

type service struct{
    carRepo repository.CarStorer
}

type Service interface{
    AddCarService(ctx context.Context , req specs.CreateCarRequest) (int64 , error)
}

func NewService(carRepo repository.CarStorer) Service{
    return &service{
        carRepo:  carRepo,
    }
}


// AddCarService is used as the service layer for adding the new car into the database
func (carService *service) AddCarService(ctx context.Context , req specs.CreateCarRequest) (int64 , error){

    carRepo := model.CarRepo{
        OwnerID: req.OwnerID,
        Model: req.Model,
        RentPerHour: req.RentPerHour,
        IsAvailable: req.IsAvailable,
        CreatedAt: time.Now().Unix(),
        UpdatedAt: time.Now().Unix(),
        Longitude: req.Longitude,
        Latitude: req.Latitude,
    }   

    carID, err := carService.carRepo.AddCarRepo(ctx , carRepo)
    if err != nil{
        zap.S().Errorw("Failed to add car", "error", err)
        return 0, err
    }   

    return carID, nil
}

