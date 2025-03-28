package validation

import (
	"encoding/json"
	"errors"

	"github.com/DaniloBelem09-0-0/ApiGoOrganizada/src/configuration/rest_err"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		un := ut.New(en, en)
		transl, _ = un.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, transl)
	}
}

func ValidatorUserError(validationErr error) *rest_err.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validationErr, &jsonErr) {
		return rest_err.NewBadRequestError("Invalid field type")
	}

	if errors.As(validationErr, &jsonValidationError) {
		errorsCauses := []rest_err.Causes{}

		for _, e := range validationErr.(validator.ValidationErrors) {
			errorsCauses = append(errorsCauses, rest_err.Causes{
				Message: e.Translate(transl),
				Field:   e.Field(),
			})
		}

		return rest_err.NewBadRequestValidationError("Some fields are invalid", errorsCauses)
	}

	return rest_err.NewBadRequestError("Error trying to convert fields")
}