package middleware

import (
	"github.com/harisspace/fisheries-api/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func VerifyBasicAuth() echo.MiddlewareFunc {
	return middleware.BasicAuth(func(username, password string, context echo.Context) (bool, error) {
		if username == config.GlobalEnv.BasicAuthUsername && password == config.GlobalEnv.BasicAuthPassword {
			return true, nil
		}

		return false, nil
	})
}
