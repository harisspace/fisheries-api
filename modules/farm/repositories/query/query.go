package query

import (
	"context"

	model "github.com/harisspace/fisheries-api/modules/farm/models"
	"github.com/harisspace/fisheries-api/pkg/utils"
)

type IFarmPostgres interface {
	// Payload data type use map instead of struct for evaluated non-zero field
	// https://gorm.io/docs/query.html#specify_search_fields
	FindOne(ctx context.Context, payload map[string]interface{}) <-chan utils.Result
	FindMany(ctx context.Context, page int, quantity int, order string) <-chan utils.PaginationResult
	CountData(ctx context.Context, payload *model.Farm) <-chan utils.Result
}

type IPondPostgres interface {
	FindOne(ctx context.Context, payload map[string]interface{}) <-chan utils.Result
	FindMany(ctx context.Context, page int, quantity int, order string) <-chan utils.PaginationResult
	CountData(ctx context.Context, payload *model.Pond) <-chan utils.Result
}
