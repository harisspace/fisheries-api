package main

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/harisspace/fisheries-api/config"
	customMiddleware "github.com/harisspace/fisheries-api/middleware"
	httpHandlerFarm "github.com/harisspace/fisheries-api/modules/farm/handlers"
	httpHandlerStatistic "github.com/harisspace/fisheries-api/modules/statistic/handlers"
	"github.com/harisspace/fisheries-api/pkg/database"
	"github.com/harisspace/fisheries-api/pkg/utils"
)

func main() {
	// Setup DB
	database.InitPosgres()

	e := echo.New()

	e.Use(middleware.CORS())

	// Register validator
	echoCustom := utils.NewEchoCustom()
	e.Validator = echoCustom

	// Instance of handler
	farmHTTP := httpHandlerFarm.NewFarmHandler()
	statisticHTTP := httpHandlerStatistic.NewStatisticHandler()

	// Mount http handler group
	farmHTTP.MountFarm(e.Group("/v1/farm", customMiddleware.RecordApi))
	farmHTTP.MountPond(e.Group("/v1/pond", customMiddleware.RecordApi))
	statisticHTTP.MountStatistic(e.Group("/v1/statistic"))

	e.Logger.Fatal(e.Start(":" + strconv.Itoa(int(config.GlobalEnv.Port))))
}
