package specs

import (
	"errors"
	"strconv"
	"strings"
)



type CreateCarRequest struct {
	OwnerID     int64    `json:"owner_id"`      
	Model       string   `json:"model"`         
	RentPerHour float64  `json:"rent_per_hour"`
	IsAvailable bool     `json:"is_available"` 
	Latitude    string 	`json:"latitude"`      
	Longitude  string  `json:"longitude"` 
}

func (r *CreateCarRequest) Validate() error {
	if r.OwnerID <= 0 {
		return errors.New("owner_id cannot be empty or less than 1")
	}

	if strings.TrimSpace(r.Model) == "" {
		return errors.New("model cannot be empty")
	}

	if r.RentPerHour <= 0 {
		return errors.New("rent_per_hour must be greater than 0")
	}

		lat, err := strconv.ParseFloat(r.Latitude, 64)
	if err != nil {
		return errors.New("latitude must be a valid number")
	}
	if lat < -90 || lat > 90 {
		return errors.New("latitude must be between -90 and 90")
	}

	// Convert longitude to float64 and validate the range
	long, err := strconv.ParseFloat(r.Longitude, 64)
	if err != nil {
		return errors.New("longitude must be a valid number")
	}
	if long < -180 || long > 180 {
		return errors.New("longitude must be between -180 and 180")
	}

	// Validation for IsAvailable
	if !r.IsAvailable && r.IsAvailable {
		return errors.New("is_available must be explicitly set to true or false")
	}

	return nil
}




