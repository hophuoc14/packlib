package validation

import (
	"packlib/entity"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func ValidateDepartment(param entity.Department, isCreate bool) error {
	if !isCreate && param.Name == "" && param.Description == "" {
		return validation.NewError("missing_fields", "at least one field must be provided")
	}
	
	nameValidation := validation.Required
	if !isCreate {
		nameValidation = validation.NilOrNotEmpty
	}

	descriptionValidation := validation.Required
	if !isCreate {
		descriptionValidation = validation.NilOrNotEmpty
	}
	error := validation.ValidateStruct(&param,
		validation.Field(&param.Name, nameValidation),
		validation.Field(&param.Description, descriptionValidation),
	)

	if error != nil {
		return error
	}

	return nil
}
