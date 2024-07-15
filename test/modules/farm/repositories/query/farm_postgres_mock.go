package query_test

import (
	"context"

	model "github.com/harisspace/fisheries-api/modules/farm/models"
	"github.com/harisspace/fisheries-api/pkg/utils"
	"github.com/harisspace/fisheries-api/test/data"

	"github.com/stretchr/testify/mock"
)

type FarmQueryMock struct {
	Mock mock.Mock
}

func (q *FarmQueryMock) FindOne(ctx context.Context, payload map[string]interface{}) <-chan utils.Result {
	output := make(chan utils.Result)

	arguments := q.Mock.Called(ctx, payload)
	go func() {
		defer close(output)

		if arguments.Get(0).(utils.Result).Error != nil {
			output <- utils.Result{Error: "Error findOne"}
		} else if arguments.Get(0).(utils.Result).Data == nil {
			output <- utils.Result{Data: nil}
		}

		output <- utils.Result{Data: data.SingleFarm}
	}()

	return output
}

func (q *FarmQueryMock) FindMany(ctx context.Context, page int, quantity int, order string) <-chan utils.PaginationResult {
	output := make(chan utils.PaginationResult)

	arguments := q.Mock.Called(ctx, page, quantity, order)
	go func() {
		defer close(output)

		if arguments.Get(0).(utils.PaginationResult).Error != nil {
			output <- utils.PaginationResult{Error: "Error not found"}
		}

		output <- utils.PaginationResult{Data: data.ManyFarm}
	}()

	return output
}

func (q *FarmQueryMock) CountData(ctx context.Context, farm *model.Farm) <-chan utils.Result {
	output := make(chan utils.Result)

	arguments := q.Mock.Called(ctx, farm)
	go func() {
		defer close(output)

		if arguments.Get(0).(utils.Result).Error != nil {
			output <- utils.Result{Error: "Error count data"}
		}
	}()

	output <- utils.Result{Data: 5}

	return output
}
