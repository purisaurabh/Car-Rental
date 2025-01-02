package decoder

import (
	"encoding/json"
	"net/http"

	"github.com/purisaurabh/car-rental/internal/pkg/specs"
	"go.uber.org/zap"
)

func DecodeAddCarRequest(r *http.Request)(specs.CreateCarRequest , error){
	var req specs.CreateCarRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		zap.S().Error("Error while decoding the car registration request : ", err)
		return specs.CreateCarRequest{}, err
	}
	return req, nil
}
