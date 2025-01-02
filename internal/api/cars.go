package api

import (
	"context"
	"net/http"

	"github.com/purisaurabh/car-rental/internal/api/decoder"
	"github.com/purisaurabh/car-rental/internal/app/car"
	specs "github.com/purisaurabh/car-rental/internal/pkg/responses"
)

func CreateCar(ctx context.Context, carService car.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		req, err := decoder.DecodeAddCarRequest(r)
		if err != nil {
			specs.HandleError(w, http.StatusBadRequest, "Failed to decode car registration request", err)
			return
		}

		// validation
		if err := req.Validate(); err != nil {
			specs.HandleError(w, http.StatusBadRequest, "Invalid car registration request", err)
			return
		}

		res, err := carService.AddCarService(ctx, req)
		if err != nil {
			specs.HandleError(w, http.StatusInternalServerError, "Failed to register car", err)
			return
		}

		specs.SuccessResponse(w, http.StatusCreated, specs.MessageResponseWithUserID{
			ID:      res,
			Message: "Car registered successfully",
		})

	}
}
