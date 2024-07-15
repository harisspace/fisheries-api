package command_test

import (
	"context"

	model "github.com/harisspace/fisheries-api/modules/farm/models"
	"github.com/harisspace/fisheries-api/pkg/utils"
	"github.com/harisspace/fisheries-api/test/data"

	"github.com/stretchr/testify/mock"
)

type PondCommandMock struct {
	Mock mock.Mock
}

func (c *PondCommandMock) InsertOne(ctx context.Context, value *model.Pond) <-chan utils.Result {
	output := make(chan utils.Result)

	arguments := c.Mock.Called(ctx, value)
	go func() {
		defer close(output)

		if arguments.Get(0).(utils.Result).Error != nil {
			output <- utils.Result{Error: "Error findOne"}
		}

		output <- utils.Result{Data: data.SinglePond}
	}()

	return output
}

func (c *PondCommandMock) UpdateOne(ctx context.Context, payload map[string]interface{}, value map[string]interface{}) <-chan utils.Result {
	output := make(chan utils.Result)

	arguments := c.Mock.Called(ctx, payload, value)
	go func() {
		defer close(output)

		if arguments.Get(0).(utils.Result).Error != nil {
			output <- utils.Result{Error: "Update failed"}
		}

		output <- utils.Result{Data: data.SinglePond}
	}()

	return output
}
