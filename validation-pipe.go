package scarlet

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type ValidationPipe struct {
}

func ValidateBody[T any](dto T) ScarletRouteHandler {
	validate := validator.New()

	return func(ctx ScarletRequestContext) interface{} {
		var body T

		err := json.NewDecoder(ctx.Request.Body).Decode(&body)

		if err != nil {
			return ScarletError{
				StatusCode: http.StatusBadRequest,
				Message:    "Invalid body",
			}
		}

		if err := validate.Struct(body); err != nil {
			return processValidationErrors(err)
		}

		return nil
	}
}

func processValidationErrors(err error) ScarletError {
	errorsMap := make(map[string]string)

	for _, err := range err.(validator.ValidationErrors) {
		errorsMap[err.Field()] = err.Error()
	}

	jsonErrors, _ := json.Marshal(errorsMap)

	return ScarletError{
		StatusCode: http.StatusBadRequest,
		Message:    string(jsonErrors),
	}
}
