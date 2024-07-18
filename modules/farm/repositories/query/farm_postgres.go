package query

import (
	"context"
	"fmt"
	"math"

	model "github.com/harisspace/fisheries-api/modules/farm/models"
	"github.com/harisspace/fisheries-api/pkg/utils"
	"gorm.io/gorm"
)

type FarmPostgres struct {
	db *gorm.DB
}

func NewFarmQueryPostgres(db *gorm.DB) *FarmPostgres {
	return &FarmPostgres{
		db: db,
	}
}

func (q *FarmPostgres) FindOne(ctx context.Context, payload map[string]interface{}) <-chan utils.Result {
	output := make(chan utils.Result)

	go func() {
		defer close(output)

		var farm model.Farm

		result := q.db.Preload("Pond").Model(&farm).Where(payload).First(&farm)
		if result.Error != nil {
			output <- utils.Result{Error: result}
			return
		}

		output <- utils.Result{Data: farm}
	}()

	return output
}

func (q *FarmPostgres) FindMany(ctx context.Context, page int, quantity int, order string) <-chan utils.PaginationResult {
	output := make(chan utils.PaginationResult)

	go func() {
		defer close(output)

		var farm []model.Farm

		offset := (page - 1) * quantity

		result := q.db.Limit(quantity).Offset(offset).Order(fmt.Sprintf("%s %s", "created_at", order)).Where("is_deleted=?", false).Find(&farm)
		if result.Error != nil || result.RowsAffected == 0 {
			output <- utils.PaginationResult{Error: result}
			return
		}
		countRows := <-q.CountData(ctx, &model.Farm{})

		paginationResult := utils.PaginationResult{
			Data: farm,
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

func (q *FarmPostgres) CountData(ctx context.Context, farm *model.Farm) <-chan utils.Result {
	output := make(chan utils.Result)

	go func() {
		var count int64 = 0
		q.db.Model(&farm).Where("is_deleted=?", false).Count(&count)

		countResult := utils.Result{
			Data: count,
		}

		output <- countResult
	}()

	return output
}
