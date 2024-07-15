package usecase_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	model "github.com/harisspace/fisheries-api/modules/farm/models"
	usecase "github.com/harisspace/fisheries-api/modules/farm/usecases"
	httpError "github.com/harisspace/fisheries-api/pkg/http_error"
	"github.com/harisspace/fisheries-api/pkg/utils"
	"github.com/harisspace/fisheries-api/test/data"
	queryTest "github.com/harisspace/fisheries-api/test/modules/farm/repositories/query"
)

var (
	farmQuery                = queryTest.FarmQueryMock{Mock: mock.Mock{}}
	pondQuery                = queryTest.PondQueryMock{Mock: mock.Mock{}}
	farmQueryUsecasePostgres = usecase.NewFarmQueryUsecasePostgres(&farmQuery, &pondQuery)
)

func TestFarmQueryUsecase_GetManyFarm(t *testing.T) {
	t.Run("GetManyFarm_Positive#1", func(t *testing.T) {
		payload := model.GetManyFarm{
			GetMany: model.GetMany{
				Page:     1,
				Quantity: 10,
				Order:    "desc",
			},
		}
		// Mock function
		mockCall := farmQuery.Mock.On("FindMany", context.Background(), payload.Page, payload.Quantity, payload.Order).Return(utils.PaginationResult{Data: data.ManyFarm})

		// Test usecase
		result := farmQueryUsecasePostgres.GetManyFarm(context.Background(), &payload)

		assert.Equal(t, data.ManyFarm, result.Data)

		mockCall.Unset()
	})

	t.Run("GetManyFarm_Negative#1", func(t *testing.T) {
		payload := model.GetManyFarm{
			GetMany: model.GetMany{
				Page:     1,
				Quantity: 10,
				Order:    "desc",
			},
		}

		// Error obj
		errObj := httpError.NewNotFound()
		errObj.Message = "Farm not found"
		mockCall := farmQuery.Mock.On("FindMany", context.Background(), payload.Page, payload.Quantity, payload.Order).Return(utils.PaginationResult{Error: errObj})

		result := farmQueryUsecasePostgres.GetManyFarm(context.Background(), &payload)

		assert.Equal(t, errObj, result.Error)

		mockCall.Unset()
	})
}

func TestFarmQueryUsecase_GetSingleFarmById(t *testing.T) {
	t.Run("GetSingleFarmById_Positive#1", func(t *testing.T) {
		payload := model.GetFarmById{
			FarmId: "farmId",
		}
		queryParam := map[string]interface{}{"farm_id": payload.FarmId, "is_deleted": false}
		// Mock function
		mockCall := farmQuery.Mock.On("FindOne", context.Background(), queryParam).Return(utils.Result{Data: data.SingleFarm})

		// Test usecase
		result := farmQueryUsecasePostgres.GetSingleFarmById(context.Background(), &payload)

		assert.Equal(t, data.SingleFarm, result.Data)

		mockCall.Unset()
	})

	t.Run("GetSingleFarmById_Negative#1", func(t *testing.T) {
		payload := model.GetFarmById{
			FarmId: "farmId",
		}
		queryParam := map[string]interface{}{"farm_id": payload.FarmId, "is_deleted": false}

		// Error obj
		errObj := httpError.NewNotFound()
		errObj.Message = "Farm not found"
		mockCall := farmQuery.Mock.On("FindOne", context.Background(), queryParam).Return(utils.Result{Error: errObj})

		result := farmQueryUsecasePostgres.GetSingleFarmById(context.Background(), &payload)

		assert.Equal(t, errObj, result.Error)

		mockCall.Unset()
	})
}

// ================== POND USECASE ====================
func TestPondQueryUsecase_GetManyPond(t *testing.T) {
	t.Run("GetManyPond_Positive#1", func(t *testing.T) {
		payload := model.GetManyPond{
			GetMany: model.GetMany{
				Page:     1,
				Quantity: 10,
				Order:    "desc",
			},
		}
		// Mock function
		mockCall := pondQuery.Mock.On("FindMany", context.Background(), payload.Page, payload.Quantity, payload.Order).Return(utils.PaginationResult{Data: data.ManyPond})

		// Test usecase
		result := farmQueryUsecasePostgres.GetManyPond(context.Background(), &payload)

		assert.Equal(t, data.ManyPond, result.Data)

		mockCall.Unset()
	})

	t.Run("GetManyPond_Negative#1", func(t *testing.T) {
		payload := model.GetManyPond{
			GetMany: model.GetMany{
				Page:     1,
				Quantity: 10,
				Order:    "desc",
			},
		}

		// Error obj
		errObj := httpError.NewNotFound()
		errObj.Message = "Pond not found"
		mockCall := pondQuery.Mock.On("FindMany", context.Background(), payload.Page, payload.Quantity, payload.Order).Return(utils.PaginationResult{Error: errObj})

		result := farmQueryUsecasePostgres.GetManyPond(context.Background(), &payload)

		assert.Equal(t, errObj, result.Error)

		mockCall.Unset()
	})
}

func TestPondQueryUsecase_GetSinglePondById(t *testing.T) {
	t.Run("GetSinglePondById_Positive#1", func(t *testing.T) {
		payload := model.GetPondById{
			PondId: "pondId",
		}
		queryParam := map[string]interface{}{"pond_id": payload.PondId, "is_deleted": false}
		// Mock function
		mockCall := pondQuery.Mock.On("FindOne", context.Background(), queryParam).Return(utils.Result{Data: data.SinglePond})

		// Test usecase
		result := farmQueryUsecasePostgres.GetSinglePondById(context.Background(), &payload)

		assert.Equal(t, data.SinglePond, result.Data)

		mockCall.Unset()
	})

	t.Run("GetSinglePondById_Negative#1", func(t *testing.T) {
		payload := model.GetPondById{
			PondId: "pondId",
		}
		queryParam := map[string]interface{}{"pond_id": payload.PondId, "is_deleted": false}

		// Error obj
		errObj := httpError.NewNotFound()
		errObj.Message = "Pond not found"
		mockCall := pondQuery.Mock.On("FindOne", context.Background(), queryParam).Return(utils.Result{Error: errObj})

		result := farmQueryUsecasePostgres.GetSinglePondById(context.Background(), &payload)

		assert.Equal(t, errObj, result.Error)

		mockCall.Unset()
	})
}
