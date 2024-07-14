package command

import (
	"context"

	model "github.com/harisspace/fisheries-api/modules/statistic/models"
	"github.com/harisspace/fisheries-api/pkg/utils"
)

type IStatisticPostgres interface {
	InsertOne(ctx context.Context, data *model.Statistic) <-chan utils.Result
	UpdateOne(ctx context.Context, payload map[string]interface{}, value map[string]interface{}) <-chan utils.Result
}
