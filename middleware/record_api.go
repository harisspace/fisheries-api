package middleware

import (
	model "github.com/harisspace/fisheries-api/modules/statistic/models"
	"github.com/harisspace/fisheries-api/modules/statistic/repositories/command"
	"github.com/harisspace/fisheries-api/modules/statistic/repositories/query"
	"github.com/harisspace/fisheries-api/pkg/database"
	"github.com/labstack/echo/v4"
)

func RecordApi(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := next(c); err != nil {
			c.Error(err)
		}

		statistic := &model.Statistic{
			UserAgent:   c.RealIP(),
			Method:      c.Request().Method,
			RequestPath: c.Path(),
			Count:       1,
		}

		db := database.GetDB()
		queryStatisticRepository := query.NewStatisticQueryPostgres(db)
		commandStatisticRepository := command.NewStatisticCommandPostgres(db)

		queryRes := <-queryStatisticRepository.FindOne(c.Request().Context(),
			map[string]interface{}{"user_agent": statistic.UserAgent, "method": statistic.Method, "request_path": statistic.RequestPath})

		// If not exist, create new
		if queryRes.Error != nil {
			<-commandStatisticRepository.InsertOne(c.Request().Context(), statistic)
			return nil
		}

		// If exist, update data
		count := queryRes.Data.(model.Statistic).Count + 1
		<-commandStatisticRepository.UpdateOne(c.Request().Context(),
			map[string]interface{}{"user_agent": statistic.UserAgent, "method": statistic.Method, "request_path": statistic.RequestPath},
			map[string]interface{}{"count": count})
		return nil
	}
}
