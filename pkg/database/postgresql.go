package database

import (
	"fmt"

	"github.com/harisspace/fisheries-api/config"
	model "github.com/harisspace/fisheries-api/modules/farm/models"
	modelStatistic "github.com/harisspace/fisheries-api/modules/statistic/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitPosgres() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		config.GlobalEnv.PostgresHost,
		config.GlobalEnv.PostrgresUsername,
		config.GlobalEnv.PostrgresPassword,
		config.GlobalEnv.PostgresDBName,
		config.GlobalEnv.PostgresPort,
		"disable",
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(fmt.Sprintf("Failed connect to database: %s", err.Error()))
	}

	db.AutoMigrate(&model.Farm{}, &model.Pond{}, &modelStatistic.Statistic{})

	sqlDB, err := db.DB()
	if err != nil {
		panic("Failed to create pool connection database postgres")
	}

	sqlDB.SetMaxOpenConns(100)

	DB = db
}

func GetDB() *gorm.DB {
	return DB
}
