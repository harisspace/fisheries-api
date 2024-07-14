package usecase

import (
	"context"

	model "github.com/harisspace/fisheries-api/modules/farm/models"
	"github.com/harisspace/fisheries-api/pkg/utils"
)

type CommandUsecase interface {
	CreateFarm(ctx context.Context, payload *model.CreateFarm) utils.Result
	CreatePond(ctx context.Context, payload *model.CreatePond) utils.Result
	UpdateFarm(ctx context.Context, payload *model.UpdateFarm) utils.Result
	UpdatePond(ctx context.Context, payload *model.UpdatePond) utils.Result
	SoftDeleteFarmById(ctx context.Context, payload *model.DeleteFarm) utils.Result
	SoftDeletePondById(ctx context.Context, payload *model.DeletePond) utils.Result
}

type QueryUsecase interface {
	GetManyFarm(ctx context.Context, payload *model.GetManyFarm) utils.PaginationResult
	GetManyPond(ctx context.Context, payload *model.GetManyPond) utils.PaginationResult
	GetSingleFarmById(ctx context.Context, payload *model.GetFarmById) utils.Result
	GetSinglePondById(ctx context.Context, payload *model.GetPondById) utils.Result
}
