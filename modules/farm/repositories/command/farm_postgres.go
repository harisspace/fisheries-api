package command

import (
	"context"

	"gorm.io/gorm"

	model "github.com/harisspace/fisheries-api/modules/farm/models"
	"github.com/harisspace/fisheries-api/pkg/utils"
)

type FarmCommandPostgres struct {
	db *gorm.DB
}

func NewFarmCommandPostgres(db *gorm.DB) *FarmCommandPostgres {
	return &FarmCommandPostgres{
		db: db,
	}
}

func (c *FarmCommandPostgres) InsertOne(ctx context.Context, farm *model.Farm) <-chan utils.Result {
	output := make(chan utils.Result)

	go func() {
		defer close(output)
		result := c.db.Create(&farm)
		if result.Error != nil {
			output <- utils.Result{Error: result.Error}
		}
		output <- utils.Result{Data: farm}
	}()

	return output
}

func (c *FarmCommandPostgres) UpdateOne(ctx context.Context, payload map[string]interface{}, value map[string]interface{}) <-chan utils.Result {
	output := make(chan utils.Result)

	go func() {
		defer close(output)

		var farm model.Farm

		result := c.db.Model(&farm).Where(payload).Updates(value)
		if result.Error != nil {
			output <- utils.Result{Error: result.Error}
		}
		output <- utils.Result{Data: farm}
	}()

	return output
}
