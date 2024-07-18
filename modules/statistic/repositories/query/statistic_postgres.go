package query

import (
	"context"
	"fmt"
	"math"

	model "github.com/harisspace/fisheries-api/modules/statistic/models"
	"github.com/harisspace/fisheries-api/pkg/utils"
	"gorm.io/gorm"
)

type StatisticPostgres struct {
	db *gorm.DB
}

func NewStatisticQueryPostgres(db *gorm.DB) *StatisticPostgres {
	return &StatisticPostgres{
		db: db,
	}
}

func (q *StatisticPostgres) FindOne(ctx context.Context, payload map[string]interface{}) <-chan utils.Result {
	output := make(chan utils.Result)

	go func() {
		defer close(output)

		var statistic model.Statistic

		result := q.db.Model(&statistic).Where(payload).First(&statistic)
		if result.Error != nil {
			output <- utils.Result{Error: result}
			return
		}

		output <- utils.Result{Data: statistic}
	}()

	return output
}

func (q *StatisticPostgres) FindManyByPayload(ctx context.Context, payload map[string]interface{}, page int, quantity int, order string) <-chan utils.PaginationResult {
	output := make(chan utils.PaginationResult)

	go func() {
		defer close(output)

		var statistic []model.Statistic

		offset := (page - 1) * quantity

		result := q.db.Limit(quantity).Offset(offset).Order(fmt.Sprintf("%s %s", "created_at", order)).Where(payload).Find(&statistic)
		if result.Error != nil || result.RowsAffected == 0 {
			output <- utils.PaginationResult{Error: result}
			return
		}
		countRows := <-q.CountData(ctx, &model.Statistic{})

		paginationResult := utils.PaginationResult{
			Data: statistic,
			Meta: utils.PaginationMeta{
				Page:      page,
				Quantity:  int(result.RowsAffected),
				TotalPage: int(math.Ceil(float64(countRows.Data.(int64)) / float64(quantity))),
				TotalData: int(countRows.Data.(int64)),
			},
		}

		output <- paginationResult
	}()

	return output
}

func (q *StatisticPostgres) CountData(ctx context.Context, statistic *model.Statistic) <-chan utils.Result {
	output := make(chan utils.Result)

	go func() {
		var count int64 = 0
		q.db.Model(&statistic).Count(&count)

		countResult := utils.Result{
			Data: count,
		}

		output <- countResult
	}()

	return output
}
