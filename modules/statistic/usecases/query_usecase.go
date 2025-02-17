package usecase

import (
	"context"

	model "github.com/harisspace/fisheries-api/modules/statistic/models"
	"github.com/harisspace/fisheries-api/modules/statistic/repositories/query"
	httpError "github.com/harisspace/fisheries-api/pkg/http_error"
	"github.com/harisspace/fisheries-api/pkg/utils"
)

type StatisticQueryUsecasePostgres struct {
	statisticQueryPostgres query.IStatisticPostgres
}

func NewStatisticQueryUsecasePostgres(statisticQueryPostgres query.IStatisticPostgres) *StatisticQueryUsecasePostgres {
	return &StatisticQueryUsecasePostgres{
		statisticQueryPostgres: statisticQueryPostgres,
	}
}

func (q *StatisticQueryUsecasePostgres) GetStatisticByUserAgent(ctx context.Context, payload *model.GetStatisticByUserAgent) utils.PaginationResult {
	var result utils.PaginationResult

	userAgent := payload.UserAgent

	// Default pagination
	page := 1
	quantity := 10
	order := "desc"

	queryRes := <-q.statisticQueryPostgres.FindManyByPayload(ctx, map[string]interface{}{"user_agent": userAgent}, page, quantity, order)
	if queryRes.Error != nil {
		errObj := httpError.NewNotFound()
		errObj.Message = "User-agent not found"
		result.Error = errObj
		return result
	}

	statisticResponse := make([]map[string]model.Statistic, 0, 10)
	for _, v := range queryRes.Data.([]model.Statistic) {
		strMethodPath := v.Method + " " + v.RequestPath
		newVal := map[string]model.Statistic{
			strMethodPath: v,
		}

		statisticResponse = append(statisticResponse, newVal)
	}

	queryRes.Data = statisticResponse

	result = queryRes
	return result
}
