package handler

import (
	"net/http"

	"github.com/harisspace/fisheries-api/middleware"
	model "github.com/harisspace/fisheries-api/modules/farm/models"
	"github.com/harisspace/fisheries-api/modules/farm/repositories/command"
	"github.com/harisspace/fisheries-api/modules/farm/repositories/query"
	usecase "github.com/harisspace/fisheries-api/modules/farm/usecases"
	"github.com/harisspace/fisheries-api/pkg/database"
	httpError "github.com/harisspace/fisheries-api/pkg/http_error"
	"github.com/harisspace/fisheries-api/pkg/utils"

	"github.com/labstack/echo/v4"
)

type farmHTTPHandler struct {
	farmCommandUsecase usecase.CommandUsecase
	farmQueryUsecase   usecase.QueryUsecase
}

func NewFarmHandler() *farmHTTPHandler {
	// Get DB
	postgres := database.GetDB()
	// Repositories
	farmCommandPostgres := command.NewFarmCommandPostgres(postgres)
	farmQueryPostgres := query.NewFarmQueryPostgres(postgres)
	pondCommandPostgres := command.NewPondCommandPostgres(postgres)
	pondQueryPostgres := query.NewPondQueryPostgres(postgres)
	// Usecase
	farmCommandUsecase := usecase.NewFarmCommandUsecasePostgres(farmCommandPostgres, farmQueryPostgres, pondCommandPostgres, pondQueryPostgres)
	farmQueryUsecase := usecase.NewFarmQueryUsecasePostgres(farmQueryPostgres, pondQueryPostgres)

	return &farmHTTPHandler{
		farmCommandUsecase: farmCommandUsecase,
		farmQueryUsecase:   farmQueryUsecase,
	}
}

func (f *farmHTTPHandler) MountFarm(e *echo.Group) {
	e.POST("", f.CreateFarm, middleware.VerifyBasicAuth())
	e.GET("", f.GetManyFarm, middleware.VerifyBasicAuth())
	e.GET("/:id", f.GetFarmById, middleware.VerifyBasicAuth())
	e.PUT("", f.UpdateFarm, middleware.VerifyBasicAuth())
	e.DELETE("/:id", f.SoftDeleteFarmById, middleware.VerifyBasicAuth())
}

func (f *farmHTTPHandler) MountPond(e *echo.Group) {
	e.POST("", f.CreatePond, middleware.VerifyBasicAuth())
	e.GET("", f.GetManyPond, middleware.VerifyBasicAuth())
	e.GET("/:id", f.GetPondById, middleware.VerifyBasicAuth())
	e.PUT("", f.UpdatePond, middleware.VerifyBasicAuth())
	e.DELETE("/:id", f.SoftDeletePondById, middleware.VerifyBasicAuth())
}

// ============== FARM HTTP HANDLER METHOD ==================
func (f *farmHTTPHandler) CreateFarm(c echo.Context) error {
	var data model.CreateFarm
	err := c.Bind(&data)
	if err != nil {
		errOjb := httpError.NewBadRequest()
		errOjb.Message = err.Error()
		return utils.ResponseError(errOjb, c)
	}

	if err = c.Validate(&data); err != nil {
		return utils.ResponseError(err.(*echo.HTTPError).Message, c)
	}

	result := f.farmCommandUsecase.CreateFarm(c.Request().Context(), &data)
	if result.Error != nil {
		return utils.ResponseError(result.Error, c)
	}

	return utils.ResponseSuccess(result.Data, "Farm created", http.StatusOK, c)
}

func (f *farmHTTPHandler) GetManyFarm(c echo.Context) error {
	var data model.GetManyFarm

	err := c.Bind(&data)
	if err != nil {
		errOjb := httpError.NewBadRequest()
		errOjb.Message = err.Error()
		return utils.ResponseError(errOjb, c)
	}

	if err = c.Validate(&data); err != nil {
		return utils.ResponseError(err.(*echo.HTTPError).Message, c)
	}

	result := f.farmQueryUsecase.GetManyFarm(c.Request().Context(), &data)
	if result.Error != nil {
		return utils.ResponseError(result.Error, c)
	}

	return utils.PaginationResponseSuccess(result.Data, "", http.StatusOK, result.Meta, c)
}

func (f *farmHTTPHandler) GetFarmById(c echo.Context) error {
	var data model.GetFarmById

	err := c.Bind(&data)
	if err != nil {
		errOjb := httpError.NewBadRequest()
		errOjb.Message = err.Error()
		return utils.ResponseError(errOjb, c)
	}

	if err = c.Validate(&data); err != nil {
		return utils.ResponseError(err.(*echo.HTTPError).Message, c)
	}

	result := f.farmQueryUsecase.GetSingleFarmById(c.Request().Context(), &data)
	if result.Error != nil {
		return utils.ResponseError(result.Error, c)
	}

	return utils.ResponseSuccess(result.Data, "", http.StatusOK, c)
}

func (f *farmHTTPHandler) UpdateFarm(c echo.Context) error {
	var data model.UpdateFarm
	err := c.Bind(&data)
	if err != nil {
		errOjb := httpError.NewBadRequest()
		errOjb.Message = err.Error()
		return utils.ResponseError(errOjb, c)
	}

	if err = c.Validate(&data); err != nil {
		return utils.ResponseError(err.(*echo.HTTPError).Message, c)
	}

	result := f.farmCommandUsecase.UpdateFarm(c.Request().Context(), &data)
	if result.Error != nil {
		return utils.ResponseError(result.Error, c)
	}

	return utils.ResponseSuccess(result.Data, "Farm updated", http.StatusOK, c)
}

func (f *farmHTTPHandler) SoftDeleteFarmById(c echo.Context) error {
	var data model.DeleteFarm
	err := c.Bind(&data)
	if err != nil {
		errOjb := httpError.NewBadRequest()
		errOjb.Message = err.Error()
		return utils.ResponseError(errOjb, c)
	}

	if err = c.Validate(&data); err != nil {
		return utils.ResponseError(err.(*echo.HTTPError).Message, c)
	}

	result := f.farmCommandUsecase.SoftDeleteFarmById(c.Request().Context(), &data)
	if result.Error != nil {
		return utils.ResponseError(result.Error, c)
	}

	return utils.ResponseSuccess(result.Data, "Farm deleted", http.StatusOK, c)
}

// ============== POND HTTP HANDLER METHOD ==================
func (p *farmHTTPHandler) CreatePond(c echo.Context) error {
	var data model.CreatePond
	err := c.Bind(&data)
	if err != nil {
		errOjb := httpError.NewBadRequest()
		errOjb.Message = err.Error()
		return utils.ResponseError(errOjb, c)
	}

	if err = c.Validate(&data); err != nil {
		return utils.ResponseError(err.(*echo.HTTPError).Message, c)
	}

	result := p.farmCommandUsecase.CreatePond(c.Request().Context(), &data)
	if result.Error != nil {
		return utils.ResponseError(result.Error, c)
	}

	return utils.ResponseSuccess(result.Data, "Pond created", http.StatusOK, c)
}

func (p *farmHTTPHandler) GetManyPond(c echo.Context) error {
	var data model.GetManyPond

	err := c.Bind(&data)
	if err != nil {
		errOjb := httpError.NewBadRequest()
		errOjb.Message = err.Error()
		return utils.ResponseError(errOjb, c)
	}

	if err = c.Validate(&data); err != nil {
		return utils.ResponseError(err.(*echo.HTTPError).Message, c)
	}

	result := p.farmQueryUsecase.GetManyPond(c.Request().Context(), &data)
	if result.Error != nil {
		return utils.ResponseError(result.Error, c)
	}

	return utils.PaginationResponseSuccess(result.Data, "", http.StatusOK, result.Meta, c)
}

func (p *farmHTTPHandler) GetPondById(c echo.Context) error {
	var data model.GetPondById

	err := c.Bind(&data)
	if err != nil {
		errOjb := httpError.NewBadRequest()
		errOjb.Message = err.Error()
		return utils.ResponseError(errOjb, c)
	}

	if err = c.Validate(&data); err != nil {
		return utils.ResponseError(err.(*echo.HTTPError).Message, c)
	}

	result := p.farmQueryUsecase.GetSinglePondById(c.Request().Context(), &data)
	if result.Error != nil {
		return utils.ResponseError(result.Error, c)
	}

	return utils.ResponseSuccess(result.Data, "", http.StatusOK, c)
}

func (p *farmHTTPHandler) UpdatePond(c echo.Context) error {
	var data model.UpdatePond
	err := c.Bind(&data)
	if err != nil {
		errOjb := httpError.NewBadRequest()
		errOjb.Message = err.Error()
		return utils.ResponseError(errOjb, c)
	}

	if err = c.Validate(&data); err != nil {
		return utils.ResponseError(err.(*echo.HTTPError).Message, c)
	}

	result := p.farmCommandUsecase.UpdatePond(c.Request().Context(), &data)
	if result.Error != nil {
		return utils.ResponseError(result.Error, c)
	}

	return utils.ResponseSuccess(result.Data, "Pond updated", http.StatusOK, c)
}

func (p *farmHTTPHandler) SoftDeletePondById(c echo.Context) error {
	var data model.DeletePond
	err := c.Bind(&data)
	if err != nil {
		errOjb := httpError.NewBadRequest()
		errOjb.Message = err.Error()
		return utils.ResponseError(errOjb, c)
	}

	if err = c.Validate(&data); err != nil {
		return utils.ResponseError(err.(*echo.HTTPError).Message, c)
	}

	result := p.farmCommandUsecase.SoftDeletePondById(c.Request().Context(), &data)
	if result.Error != nil {
		return utils.ResponseError(result.Error, c)
	}

	return utils.ResponseSuccess(result.Data, "Pond deleted", http.StatusOK, c)
}
