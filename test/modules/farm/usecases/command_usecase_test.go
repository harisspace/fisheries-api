package usecase_test

// import (
// 	"github.com/stretchr/testify/mock"

// 	usecase "github.com/harisspace/fisheries-api/modules/farm/usecases"
// 	commandTest "github.com/harisspace/fisheries-api/test/modules/farm/repositories/command"
// 	queryTest "github.com/harisspace/fisheries-api/test/modules/farm/repositories/query"
// )

// var (
// 	c_farmCommand                = commandTest.FarmCommandMock{Mock: mock.Mock{}}
// 	c_pondCommand                = commandTest.PondCommandMock{Mock: mock.Mock{}}
// 	c_farmQuery                  = queryTest.FarmQueryMock{Mock: mock.Mock{}}
// 	c_pondQuery                  = queryTest.PondQueryMock{Mock: mock.Mock{}}
// 	c_farmCommandUsecasePostgres = usecase.NewFarmCommandUsecasePostgres(&c_farmCommand, &c_farmQuery, &c_pondCommand, &c_pondQuery)
// )

// func TestFarmCommandUsecase_CreateFarm(t *testing.T) {
// 	t.Run("CreateFarm_Positive#1", func(t *testing.T) {
// 		payload1 := model.CreateFarm{
// 			Name: "farm1",
// 		}
// 		payload2 := model.Farm{
// 			FarmId:    "farmId",
// 			Name:      payload1.Name,
// 			IsDeleted: false,
// 		}
// 		queryParam := map[string]interface{}{"name": payload1.Name}
// 		// Mock function
// 		mockCall1 := c_farmQuery.Mock.On("FindOne", context.Background(), queryParam).Return(utils.Result{Data: nil})
// 		mockCall2 := c_farmCommand.Mock.On("InsertOne", mock.Anything).Return(utils.Result{Data: data.SingleFarm})

// 		// Test usecase
// 		result := c_farmCommandUsecasePostgres.CreateFarm(context.Background(), &payload1)

// 		// Mapping

// 		farmResponseData := &model.FarmResponse{
// 			FarmId:    payload2.FarmId,
// 			Name:      payload2.Name,
// 			CreatedAt: time.Now(),
// 			UpdatedAt: time.Now(),
// 		}

// 		assert.Equal(t, farmResponseData, result.Data)

// 		mockCall1.Unset()
// 		mockCall2.Unset()
// 	})
// }

// func TestFarmCommandUsecase_UpdateFarm(t *testing.T) {
// 	t.Run("UpdateFarm_Positive#1", func(t *testing.T) {
// 		payload1 := model.UpdateFarm{
// 			Name:   "farmName",
// 			FarmId: "farmId",
// 		}
// 		payload2 := model.Farm{
// 			FarmId:    "farmId",
// 			Name:      payload1.Name,
// 			IsDeleted: false,
// 		}
// 		queryParam := map[string]interface{}{"farm_id": payload1.FarmId, "is_deleted": false}
// 		// Mock function
// 		mockCall1 := c_farmQuery.Mock.On("FindOne", context.Background(), queryParam).Return(utils.Result{Data: data.SingleFarm})
// 		mockCall2 := c_farmCommand.Mock.On("InsertOne", mock.Anything).Return(utils.Result{Data: data.SingleFarm})

// 		// Test usecase
// 		result := c_farmCommandUsecasePostgres.CreateFarm(context.Background(), &payload1)

// 		// Mapping

// 		farmResponseData := &model.FarmResponse{
// 			FarmId:    payload2.FarmId,
// 			Name:      payload2.Name,
// 			CreatedAt: time.Now(),
// 			UpdatedAt: time.Now(),
// 		}

// 		assert.Equal(t, farmResponseData, result.Data)

// 		mockCall1.Unset()
// 		mockCall2.Unset()
// 	})
// }
