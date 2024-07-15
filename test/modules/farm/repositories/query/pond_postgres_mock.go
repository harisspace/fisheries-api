package query_test

import (
	"context"

	model "github.com/harisspace/fisheries-api/modules/farm/models"
	"github.com/harisspace/fisheries-api/pkg/utils"
	"github.com/harisspace/fisheries-api/test/data"

	"github.com/stretchr/testify/mock"
)

type PondQueryMock struct {
	Mock mock.Mock
}

func (q *PondQueryMock) FindOne(ctx context.Context, payload map[string]interface{}) <-chan utils.Result {
	output := make(chan utils.Result)

	arguments := q.Mock.Called(ctx, payload)
	go func() {
		defer close(output)

		if arguments.Get(0).(utils.Result).Error != nil {
			output <- utils.Result{Error: "Error findOne"}
		}

		output <- utils.Result{Data: data.SinglePond}
	}()

	return output
}

func (q *PondQueryMock) FindMany(ctx context.Context, page int, quantity int, order string) <-chan utils.PaginationResult {
	output := make(chan utils.PaginationResult)

	arguments := q.Mock.Called(ctx, page, quantity, order)
	go func() {
		defer close(output)

		if arguments.Get(0).(utils.PaginationResult).Error != nil {
			output <- utils.PaginationResult{Error: "Error findOne"}
		}

		output <- utils.PaginationResult{Data: data.ManyPond}
	}()
	return output
}

func (q *PondQueryMock) CountData(ctx context.Context, payload *model.Pond) <-chan utils.Result {
	output := make(chan utils.Result)

	arguments := q.Mock.Called(ctx, payload)
	go func() {
		defer close(output)

		if arguments.Get(0).(utils.Result).Error != nil {
			output <- utils.Result{Error: "Error count data"}
		}

		output <- utils.Result{Data: 5}
	}()
	return output
}
