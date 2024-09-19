package utils

import "github.com/go-playground/validator/v10"

func ValidateData(data interface{}) ([]string, error) {
	validate := validator.New()

	err := validate.Struct(data)
	if err == nil {
		return nil, nil
	}

	var errFields []string
	for _, fieldErr := range err.(validator.ValidationErrors) {
		errFields = append(errFields, fieldErr.Field())
	}

	return errFields, err
}
