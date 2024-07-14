package usecase

import (
	"context"

	model "github.com/harisspace/fisheries-api/modules/farm/models"
	"github.com/harisspace/fisheries-api/modules/farm/repositories/query"
	httpError "github.com/harisspace/fisheries-api/pkg/http_error"
	"github.com/harisspace/fisheries-api/pkg/utils"
)

type FarmQueryUsecasePostgres struct {
	farmQueryPostgres query.IFarmPostgres
	pondQueryPostgres query.IPondPostgres
}

func NewFarmQueryUsecasePostgres(farmQueryPostgres query.IFarmPostgres, pondQueryPostgres query.IPondPostgres) *FarmQueryUsecasePostgres {
	return &FarmQueryUsecasePostgres{
		farmQueryPostgres: farmQueryPostgres,
		pondQueryPostgres: pondQueryPostgres,
	}
}

// ============== FARM ==============
func (q *FarmQueryUsecasePostgres) GetManyFarm(ctx context.Context, payload *model.GetManyFarm) utils.PaginationResult {
	var result utils.PaginationResult

	page := payload.Page
	quantity := payload.Quantity
	order := payload.Order

	// Default pagination
	if page == 0 {
		page = 1
	}
	if quantity == 0 {
		quantity = 10
	}
	if order == "" {
		order = "desc"
	}

	queryRes := <-q.farmQueryPostgres.FindMany(ctx, page, quantity, order)
	if queryRes.Error != nil {
		errObj := httpError.NewNotFound()
		errObj.Message = "Farm not found"
		result.Error = errObj
		return result
	}
	result = queryRes

	return result
}

func (q *FarmQueryUsecasePostgres) GetSingleFarmById(ctx context.Context, payload *model.GetFarmById) utils.Result {
	var result utils.Result

	farmId := payload.FarmId

	queryRes := <-q.farmQueryPostgres.FindOne(ctx, map[string]interface{}{"farm_id": farmId, "is_deleted": false})
	if queryRes.Error != nil {
		errObj := httpError.NewNotFound()
		errObj.Message = "Farm not found"
		result.Error = errObj
		return result
	}
	result.Data = queryRes.Data

	return result
}

// ============== POND ==============
func (q *FarmQueryUsecasePostgres) GetManyPond(ctx context.Context, payload *model.GetManyPond) utils.PaginationResult {
	var result utils.PaginationResult

	page := payload.Page
	quantity := payload.Quantity
	order := payload.Order

	// Default pagination
	if page == 0 {
		page = 1
	}
	if quantity == 0 {
		quantity = 10
	}
	if order == "" {
		order = "desc"
	}

	queryRes := <-q.pondQueryPostgres.FindMany(ctx, page, quantity, order)
	if queryRes.Error != nil {
		errObj := httpError.NewNotFound()
		errObj.Message = "Pond not found"
		result.Error = errObj
		return result
	}
	result = queryRes

	return result
}

func (q *FarmQueryUsecasePostgres) GetSinglePondById(ctx context.Context, payload *model.GetPondById) utils.Result {
	var result utils.Result

	pondId := payload.PondId

	queryRes := <-q.pondQueryPostgres.FindOne(ctx, map[string]interface{}{"pond_id": pondId, "is_deleted": false})
	if queryRes.Error != nil {
		errObj := httpError.NewNotFound()
		errObj.Message = "Pond not found"
		result.Error = errObj
		return result
	}
	result.Data = queryRes.Data

	return result
}
