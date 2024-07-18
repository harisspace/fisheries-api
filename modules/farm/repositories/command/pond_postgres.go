package command

import (
	"context"

	"gorm.io/gorm"

	model "github.com/harisspace/fisheries-api/modules/farm/models"
	"github.com/harisspace/fisheries-api/pkg/utils"
)

type PondCommandPostgres struct {
	db *gorm.DB
}

func NewPondCommandPostgres(db *gorm.DB) *PondCommandPostgres {
	return &PondCommandPostgres{
		db: db,
	}
}

func (c *PondCommandPostgres) InsertOne(ctx context.Context, data *model.Pond) <-chan utils.Result {
	output := make(chan utils.Result)

	go func() {
		defer close(output)
		result := c.db.Create(&data)
		if result.Error != nil {
			output <- utils.Result{Error: result.Error}
			return
		}
		output <- utils.Result{Data: data}
	}()

	return output
}

func (c *PondCommandPostgres) UpdateOne(ctx context.Context, payload map[string]interface{}, value map[string]interface{}) <-chan utils.Result {
	output := make(chan utils.Result)

	go func() {
		defer close(output)

		var pond model.Pond

		result := c.db.Model(&pond).Where(payload).Updates(value)
		if result.Error != nil {
			output <- utils.Result{Error: result.Error}
			return
		}
		output <- utils.Result{Data: pond}
	}()

	return output
}
