package validation

import (
	"packlib/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func ValidateUser(param model.CreateUserRequest) error {	
	error := validation.ValidateStruct(&param,
		validation.Field(&param.Username, validation.Required),
		validation.Field(&param.Password, validation.Required),
		validation.Field(&param.Email, validation.Required),
	)

	if error != nil {
		return error
	}

	return nil
}
