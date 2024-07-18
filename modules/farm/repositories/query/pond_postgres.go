package query

import (
	"context"
	"fmt"
	"math"

	model "github.com/harisspace/fisheries-api/modules/farm/models"
	"github.com/harisspace/fisheries-api/pkg/utils"
	"gorm.io/gorm"
)

type PondPostgres struct {
	db *gorm.DB
}

func NewPondQueryPostgres(db *gorm.DB) *PondPostgres {
	return &PondPostgres{
		db: db,
	}
}

func (q *PondPostgres) FindOne(ctx context.Context, payload map[string]interface{}) <-chan utils.Result {
	output := make(chan utils.Result)

	go func() {
		defer close(output)

		var pond model.Pond

		result := q.db.Preload("Farm").Model(&pond).Where(payload).First(&pond)
		if result.Error != nil {
			output <- utils.Result{Error: result}
			return
		}

		output <- utils.Result{Data: pond}
	}()

	return output
}

func (q *PondPostgres) FindMany(ctx context.Context, page int, quantity int, order string) <-chan utils.PaginationResult {
	output := make(chan utils.PaginationResult)

	go func() {
		defer close(output)

		var pond []model.Pond

		offset := (page - 1) * quantity

		result := q.db.Limit(quantity).Offset(offset).Order(fmt.Sprintf("%s %s", "created_at", order)).Where("is_deleted=?", false).Find(&pond)
		if result.Error != nil || result.RowsAffected == 0 {
			output <- utils.PaginationResult{Error: result}
			return
		}
		countRows := <-q.CountData(ctx, &model.Pond{})

		paginationResult := utils.PaginationResult{
			Data: pond,
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

func (q *PondPostgres) CountData(ctx context.Context, pond *model.Pond) <-chan utils.Result {
	output := make(chan utils.Result)

	go func() {
		var count int64 = 0
		q.db.Model(&pond).Where("is_deleted=?", false).Count(&count)

		countResult := utils.Result{
			Data: count,
		}

		output <- countResult
	}()

	return output
}
