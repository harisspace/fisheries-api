package usecase

import (
	"context"

	model "github.com/harisspace/fisheries-api/modules/statistic/models"
	"github.com/harisspace/fisheries-api/pkg/utils"
)

type QueryUsecase interface {
	GetStatisticByUserAgent(ctx context.Context, payload *model.GetStatisticByUserAgent) utils.PaginationResult
}
