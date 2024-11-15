package specs

import (
	"fmt"
	"regexp"

	"github.com/purisaurabh/car-rental/internal/pkg/constants"
	"github.com/purisaurabh/car-rental/internal/pkg/errors"
)

type UserRegistrationRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Mobile   string `json:"mobile_no"`
	Role     string `json:"role"`
}

func (user *UserRegistrationRequest) Validate() error {
	if user.Name == "" {
		return fmt.Errorf("%s : name ", errors.ErrParameterMissing.Error())
	}
	if user.Email == "" {
		return fmt.Errorf("%s : email ", errors.ErrParameterMissing.Error())
	}

	matchMail, _ := regexp.MatchString(constants.EmailRegex, user.Email)
	if !matchMail {
		return fmt.Errorf("%s : email ", errors.ErrInvalidFormat.Error())
	}

	if user.Password == "" {
		return fmt.Errorf("%s : password ", errors.ErrParameterMissing.Error())
	}

	matchPassword, _ := regexp.MatchString(constants.PasswordRegex, user.Password)
	if !matchPassword {
		return fmt.Errorf("%s : password - must be at least 8 characters long and contain at least one letter, one number, and one special character", errors.ErrInvalidFormat.Error())
	}

	if user.Mobile == "" {
		return fmt.Errorf("%s : mobile_no ", errors.ErrParameterMissing.Error())
	}

	matchMob, _ := regexp.MatchString(constants.MobileRegex, user.Mobile)
	if !matchMob {
		return fmt.Errorf("%s : mobile ", errors.ErrInvalidFormat.Error())
	}

	if user.Role == "" {
		return fmt.Errorf("%s : role ", errors.ErrParameterMissing.Error())
	}

	if user.Role != "renter" && user.Role != "owner" {
		return fmt.Errorf("%s : role - must be 'renter' or 'owner'", errors.ErrInvalidFormat.Error())
	}
	return nil
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (user *UserLoginRequest) Validate() error {
	if user.Email == "" {
		return fmt.Errorf("%s : email ", errors.ErrParameterMissing.Error())
	}

	if user.Password == "" {
		return fmt.Errorf("%s : password ", errors.ErrParameterMissing.Error())
	}

	return nil
}

type UserLoginResponse struct {
	UserID int64  `json:"user_id"`
	Token  string `json:"token"`
}
