package validation

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
	"github.com/thiago1cruz/crud-go/src/configurations/error_response"
)

var (
	validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		unt := ut.New(en, en)
		transl, _ = unt.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUserError(err error) *error_response.ErrorResponse {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(err, &jsonErr) {
		return error_response.NewBadRequestError("Invalid field type")
	} else if errors.As(err, &jsonValidationError) {
		errorCauses := []error_response.Causes{}

		for _, e := range err.(validator.ValidationErrors) {
			cause := error_response.Causes{
				Field:   e.Field(),
				Message: e.Translate(transl),
			}

			errorCauses = append(errorCauses, cause)

		}

		return error_response.NewBadRequestValidationError("Some fields are invalid", errorCauses)
	} else {
		return error_response.NewBadRequestError("Error trying to convert fields")
	}
}
