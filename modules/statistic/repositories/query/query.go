package query

import (
	"context"

	"github.com/harisspace/fisheries-api/pkg/utils"
)

type IStatisticPostgres interface {
	FindManyByPayload(ctx context.Context, payload map[string]interface{}, page int, quantity int, order string) <-chan utils.PaginationResult
	FindOne(ctx context.Context, payload map[string]interface{}) <-chan utils.Result
}
