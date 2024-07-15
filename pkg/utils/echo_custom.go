package utils

import (
	"fmt"

	httpError "github.com/harisspace/fisheries-api/pkg/http_error"
	"github.com/labstack/echo/v4"

	"github.com/go-playground/validator/v10"
)

type EchoCustom struct {
	Validator *validator.Validate
}

func NewEchoCustom() *EchoCustom {
	return &EchoCustom{
		Validator: validator.New(),
	}
}

type CustomValidator struct {
	validate *validator.Validate
}

func (cv *EchoCustom) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		strErr := ""

		for _, err := range err.(validator.ValidationErrors) {
			strErr += fmt.Sprintf("%s is %s,", err.Field(), err.Tag())
		}

		errObj := httpError.NewBadRequest()
		errObj.Message = strErr

		return echo.NewHTTPError(errObj.Code, errObj)
	}
	return nil
}
