package addon

import (
	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

// API Doc: https://godoc.org/github.com/go-playground/validator
func AddValidator(server *echo.Echo) {
	server.Validator = &CustomValidator{validator: validator.New()}
}
