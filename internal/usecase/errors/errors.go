package errors

import (
	"os"

	"github.com/0x16F/cloud/users/pkg/logger"
	"github.com/goccy/go-json"
)

type Errors struct {
	errors map[int]customError
}

func New(log logger.Logger, errorsPath string) Errors {
	log = log.WithFields(logger.Fields{
		"module": "errors",
	})

	service := Errors{
		errors: make(map[int]customError),
	}

	data, err := os.ReadFile(errorsPath)
	if err != nil {
		log.Errorf("failed to read errors file: %v", err)

		return service
	}

	var errors []customError
	if err := json.Unmarshal(data, &errors); err != nil {
		log.Errorf("failed to unmarshal errors: %v", err)

		return service
	}

	for _, err := range errors {
		service.errors[err.Code] = err
	}

	return service
}

type customError struct {
	Code        int    `json:"code"`
	Message     string `json:"message"`
	Description string `json:"description"`
}

func (ce customError) Error() string {
	encoded, _ := json.Marshal(ce)
	return string(encoded)
}

func (e Errors) GetError(code int) error {
	if err, ok := e.errors[code]; ok {
		return err
	}

	return customError{
		Code:    0,
		Message: "Unknown error",
	}
}
