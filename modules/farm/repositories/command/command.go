package command

import (
	"context"

	model "github.com/harisspace/fisheries-api/modules/farm/models"
	"github.com/harisspace/fisheries-api/pkg/utils"
)

type IFarmPostgres interface {
	InsertOne(ctx context.Context, data *model.Farm) <-chan utils.Result
	UpdateOne(ctx context.Context, payload map[string]interface{}, value map[string]interface{}) <-chan utils.Result
}

type IPondPostgres interface {
	InsertOne(ctx context.Context, data *model.Pond) <-chan utils.Result
	UpdateOne(ctx context.Context, payload map[string]interface{}, value map[string]interface{}) <-chan utils.Result
}
