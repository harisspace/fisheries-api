package command

import (
	"context"

	"gorm.io/gorm"

	model "github.com/harisspace/fisheries-api/modules/statistic/models"
	"github.com/harisspace/fisheries-api/pkg/utils"
)

type StatisticCommandPostgres struct {
	db *gorm.DB
}

func NewStatisticCommandPostgres(db *gorm.DB) *StatisticCommandPostgres {
	return &StatisticCommandPostgres{
		db: db,
	}
}

func (c *StatisticCommandPostgres) InsertOne(ctx context.Context, statistic *model.Statistic) <-chan utils.Result {
	output := make(chan utils.Result)

	go func() {
		defer close(output)
		result := c.db.Create(&statistic)
		if result.Error != nil {
			output <- utils.Result{Error: result.Error}
		}
		output <- utils.Result{Data: statistic}
	}()

	return output
}

func (c *StatisticCommandPostgres) UpdateOne(ctx context.Context, payload map[string]interface{}, value map[string]interface{}) <-chan utils.Result {
	output := make(chan utils.Result)

	go func() {
		defer close(output)

		var statistic model.Statistic

		result := c.db.Model(&statistic).Where(payload).Updates(value)
		if result.Error != nil {
			output <- utils.Result{Error: result.Error}
		}
		output <- utils.Result{Data: statistic}
	}()

	return output
}
