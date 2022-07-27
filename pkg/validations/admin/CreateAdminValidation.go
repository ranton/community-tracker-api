package validations

import (
	admin "github.com/VncntDzn/community-tracker-api/pkg/admin/requests"
	"github.com/VncntDzn/community-tracker-api/pkg/validations"
	"github.com/go-playground/validator/v10"
)

func ValidateCreateAdmin(request admin.CreateAdminRequest) []*validations.ValidationErrors {
	var errors []*validations.ValidationErrors
	var validate = validator.New()
	err := validate.Struct(request)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element validations.ValidationErrors
			element.Key = err.Field()
			element.ErrorType = err.Tag()
			errors = append(errors, &element)
		}
		return errors
	}
	return nil
}
