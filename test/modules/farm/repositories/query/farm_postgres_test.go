package query

import (
	"context"

	"github.com/harisspace/fisheries-api/pkg/utils"
	"github.com/stretchr/testify/mock"
)

type FarmQueryMock struct {
	Mock mock.Mock
}

func (q *FarmQueryMock) FindOne(ctx context.Context, payload map[string]interface{}) utils.Result {
	arguments := q.Mock.Called(ctx, payload)
	if arguments.Get(0).(utils.Result).Error != nil {
		return utils.Result{Error: "Error findOne"}
	}

	return arguments.Get(0).(utils.Result)
}
