package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	model "github.com/harisspace/fisheries-api/modules/farm/models"
	"github.com/harisspace/fisheries-api/modules/farm/repositories/command"
	"github.com/harisspace/fisheries-api/modules/farm/repositories/query"
	httpError "github.com/harisspace/fisheries-api/pkg/http_error"
	"github.com/harisspace/fisheries-api/pkg/utils"
)

type FarmCommandUsecasePostgres struct {
	farmCommandPostgres command.IFarmPostgres
	pondCommandPostgres command.IPondPostgres
	farmQueryPostgres   query.IFarmPostgres
	pondQueryPostgres   query.IPondPostgres
}

func NewFarmCommandUsecasePostgres(farmCommandPostgres command.IFarmPostgres, farmQueryPostgres query.IFarmPostgres, pondCommandPostgres command.IPondPostgres, pondQueryPostgres query.IPondPostgres) *FarmCommandUsecasePostgres {
	return &FarmCommandUsecasePostgres{
		farmCommandPostgres: farmCommandPostgres,
		farmQueryPostgres:   farmQueryPostgres,
		pondCommandPostgres: pondCommandPostgres,
		pondQueryPostgres:   pondQueryPostgres,
	}
}

func (c *FarmCommandUsecasePostgres) CreateFarm(ctx context.Context, payload *model.CreateFarm) utils.Result {
	var result utils.Result

	name := payload.Name

	queryRes := <-c.farmQueryPostgres.FindOne(ctx, map[string]interface{}{"name": name})
	if queryRes.Data != nil {
		errObj := httpError.NewConflict()
		errObj.Message = "Farm with that name already exist"
		result.Error = errObj
		return result
	}

	farm := &model.Farm{
		Name:      name,
		FarmId:    uuid.New().String(),
		IsDeleted: false,
	}

	result = <-c.farmCommandPostgres.InsertOne(ctx, farm)
	if result.Error != nil {
		errObj := httpError.NewConflict()
		errObj.Message = "Failed insert one"
		result.Error = errObj
		return result
	}

	farm = result.Data.(*model.Farm)
	createFarmData := &model.FarmResponse{
		FarmId:    farm.FarmId,
		Name:      farm.Name,
		CreatedAt: farm.CreatedAt,
		UpdatedAt: farm.UpdatedAt,
	}

	result.Data = createFarmData
	return result
}

func (c *FarmCommandUsecasePostgres) UpdateFarm(ctx context.Context, payload *model.UpdateFarm) utils.Result {
	var result utils.Result

	name := payload.Name
	farmId := payload.FarmId

	queryRes := <-c.farmQueryPostgres.FindOne(ctx, map[string]interface{}{"farm_id": farmId, "is_deleted": false})
	if queryRes.Data != nil {
		// Check is farm with that name already exist
		queryRes = <-c.farmQueryPostgres.FindOne(ctx, map[string]interface{}{"name": name, "is_deleted": false})
		if queryRes.Data != nil {
			errObj := httpError.NewConflict()
			errObj.Message = "Farm with that name already exist"
			result.Error = errObj
			return result
		}
		// Update farm
		queryRes = <-c.farmCommandPostgres.UpdateOne(ctx, map[string]interface{}{"farm_id": farmId}, map[string]interface{}{"name": name, "updated_at": time.Now()})
		if queryRes.Error != nil {
			errObj := httpError.NewConflict()
			errObj.Message = "Update failed"
			result.Error = errObj
			return result
		}
		result.Data = queryRes.Data
		return result
	}

	return c.CreateFarm(ctx, &model.CreateFarm{Name: name})
}

func (c *FarmCommandUsecasePostgres) SoftDeleteFarmById(ctx context.Context, payload *model.DeleteFarm) utils.Result {
	var result utils.Result

	farmId := payload.FarmId

	queryRes := <-c.farmQueryPostgres.FindOne(ctx, map[string]interface{}{"farm_id": farmId, "is_deleted": false})
	if queryRes.Error != nil {
		errObj := httpError.NewNotFound()
		errObj.Message = "Farm not found"
		result.Error = errObj
		return result
	}

	queryRes = <-c.farmCommandPostgres.UpdateOne(ctx,
		map[string]interface{}{"farm_id": farmId},
		map[string]interface{}{"is_deleted": true, "deleted_at": time.Now()})

	if queryRes.Error != nil {
		errObj := httpError.NewConflict()
		errObj.Message = "Update failed"
		result.Error = errObj
		return result
	}

	result.Data = queryRes.Data
	return result
}

// ============== POND ==============
func (c *FarmCommandUsecasePostgres) CreatePond(ctx context.Context, payload *model.CreatePond) utils.Result {
	var result utils.Result

	name := payload.Name
	farmId := payload.FarmId

	queryRes := <-c.farmQueryPostgres.FindOne(ctx, map[string]interface{}{"farm_id": farmId})
	if queryRes.Data == nil {
		errObj := httpError.NewConflict()
		errObj.Message = "Farm with that id doesn't exist"
		result.Error = errObj
		return result
	}

	// Check pond name is already exist
	queryRes = <-c.pondQueryPostgres.FindOne(ctx, map[string]interface{}{"name": name})
	if queryRes.Data != nil {
		errObj := httpError.NewConflict()
		errObj.Message = "Pond with that name already exist"
		result.Error = errObj
		return result
	}

	pond := &model.Pond{
		Name:      name,
		FarmId:    farmId,
		PondId:    uuid.New().String(),
		IsDeleted: false,
	}

	result = <-c.pondCommandPostgres.InsertOne(ctx, pond)
	if result.Error != nil {
		errObj := httpError.NewConflict()
		errObj.Message = "Failed insert one"
		result.Error = errObj
		return result
	}

	pond = result.Data.(*model.Pond)
	pondResponse := &model.PondResponse{
		ID:        pond.ID,
		FarmId:    pond.FarmId,
		Name:      pond.Name,
		CreatedAt: pond.CreatedAt,
		UpdatedAt: pond.UpdatedAt,
	}

	result.Data = pondResponse
	return result
}

func (c *FarmCommandUsecasePostgres) UpdatePond(ctx context.Context, payload *model.UpdatePond) utils.Result {
	var result utils.Result

	name := payload.Name
	pondId := payload.PondId
	farmId := payload.FarmId

	queryRes := <-c.pondQueryPostgres.FindOne(ctx, map[string]interface{}{"pond_id": pondId})
	if queryRes.Data != nil {
		// Check is pond name exist
		queryRes = <-c.pondQueryPostgres.FindOne(ctx, map[string]interface{}{"name": name})
		if queryRes.Data != nil {
			errObj := httpError.NewConflict()
			errObj.Message = "Pond with that name already exist"
			result.Error = errObj
			return result
		}
		// Updata pond
		queryRes = <-c.pondCommandPostgres.UpdateOne(ctx, map[string]interface{}{"pond_id": pondId}, map[string]interface{}{"name": name, "updated_at": time.Now()})
		if queryRes.Error != nil {
			errObj := httpError.NewConflict()
			errObj.Message = "Update failed"
			result.Error = errObj
			return result
		}
		result.Data = queryRes.Data
		return result
	}

	if farmId != "" {
		// Create new pond
		return c.CreatePond(ctx, &model.CreatePond{Name: name, FarmId: farmId})
	}

	errObj := httpError.NewBadRequest()
	errObj.Message = "pond_id farm_id doesn't exist"
	result.Error = errObj
	return result
}

func (c *FarmCommandUsecasePostgres) SoftDeletePondById(ctx context.Context, payload *model.DeletePond) utils.Result {
	var result utils.Result

	pondId := payload.PondId

	queryRes := <-c.pondQueryPostgres.FindOne(ctx, map[string]interface{}{"pond_id": pondId, "is_deleted": false})
	if queryRes.Error != nil {
		errObj := httpError.NewNotFound()
		errObj.Message = "Pond not found"
		result.Error = errObj
		return result
	}

	queryRes = <-c.pondCommandPostgres.UpdateOne(ctx,
		map[string]interface{}{"pond_id": pondId},
		map[string]interface{}{"is_deleted": true, "deleted_at": time.Now()})

	if queryRes.Error != nil {
		errObj := httpError.NewConflict()
		errObj.Message = "Update failed"
		result.Error = errObj
		return result
	}

	result.Data = queryRes.Data
	return result
}
