package main

import (
	"fmt"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/harisspace/fisheries-api/config"
	customMiddleware "github.com/harisspace/fisheries-api/middleware"
	httpHandlerFarm "github.com/harisspace/fisheries-api/modules/farm/handlers"
	httpHandlerStatistic "github.com/harisspace/fisheries-api/modules/statistic/handlers"
	"github.com/harisspace/fisheries-api/pkg/database"
	httpError "github.com/harisspace/fisheries-api/pkg/http_error"
)

type CustomValidator struct {
	validate *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validate.Struct(i); err != nil {
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

func main() {
	// Setup DB
	database.InitPosgres()

	e := echo.New()

	e.Use(middleware.CORS())

	e.Validator = &CustomValidator{validate: validator.New()}

	farmHTTP := httpHandlerFarm.NewFarmHandler()
	statisticHTTP := httpHandlerStatistic.NewStatisticHandler()

	farmHTTP.MountFarm(e.Group("/v1/farm", customMiddleware.RecordApi))
	farmHTTP.MountPond(e.Group("/v1/pond", customMiddleware.RecordApi))
	statisticHTTP.MountStatistic(e.Group("/v1/statistic"))

	e.Logger.Fatal(e.Start(":" + strconv.Itoa(int(config.GlobalEnv.Port))))
}
