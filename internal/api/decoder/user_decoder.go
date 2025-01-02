package decoder

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/purisaurabh/car-rental/internal/pkg/specs"
)

func DecodeUserRegistrationRequest(r *http.Request) (specs.UserRegistrationRequest, error) {
	var req specs.UserRegistrationRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Println("Error while decoding the user registration request : ", err)
		return specs.UserRegistrationRequest{}, err
	}
	return req, nil
}

func DecodeUserLoginRequest(r *http.Request) (specs.UserLoginRequest, error) {
	var req specs.UserLoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Println("Error while decoding the user login request : ", err)
		return specs.UserLoginRequest{}, err
	}
	return req, nil
}
