package handler

import (
	"net/http"

	"github.com/harisspace/fisheries-api/middleware"
	model "github.com/harisspace/fisheries-api/modules/statistic/models"
	"github.com/harisspace/fisheries-api/modules/statistic/repositories/query"
	usecase "github.com/harisspace/fisheries-api/modules/statistic/usecases"
	"github.com/harisspace/fisheries-api/pkg/database"
	httpError "github.com/harisspace/fisheries-api/pkg/http_error"
	"github.com/harisspace/fisheries-api/pkg/utils"

	"github.com/labstack/echo/v4"
)

type statisticHTTPHandler struct {
	statisticQueryUsecase usecase.QueryUsecase
}

func NewStatisticHandler() *statisticHTTPHandler {
	// DB initialize
	postgres := database.GetDB()
	// Repositories
	statisticQueryPostgres := query.NewStatisticQueryPostgres(postgres)
	// Usecase
	statisticQueryUsecase := usecase.NewStatisticQueryUsecasePostgres(statisticQueryPostgres)

	return &statisticHTTPHandler{
		statisticQueryUsecase: statisticQueryUsecase,
	}
}

func (s *statisticHTTPHandler) MountStatistic(e *echo.Group) {
	e.GET("/:id", s.GetStatisticByUserAgent, middleware.VerifyBasicAuth())
}

func (s *statisticHTTPHandler) GetStatisticByUserAgent(c echo.Context) error {
	var data model.GetStatisticByUserAgent

	err := c.Bind(&data)
	if err != nil {
		errOjb := httpError.NewBadRequest()
		errOjb.Message = err.Error()
		return utils.ResponseError(errOjb, c)
	}

	if err = c.Validate(&data); err != nil {
		return utils.ResponseError(err.(*echo.HTTPError).Message, c)
	}

	result := s.statisticQueryUsecase.GetStatisticByUserAgent(c.Request().Context(), &data)
	if result.Error != nil {
		return utils.ResponseError(result.Error, c)
	}

	return utils.PaginationResponseSuccess(result.Data, "", http.StatusOK, result.Meta, c)
}
