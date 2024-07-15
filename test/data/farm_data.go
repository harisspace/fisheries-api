package data

import (
	"time"

	model "github.com/harisspace/fisheries-api/modules/farm/models"
)

var SingleFarm = model.Farm{
	ID:        1,
	FarmId:    "farmId",
	Name:      "Susi Farm",
	IsDeleted: false,
	CreatedAt: time.Now(),
	DeletedAt: time.Now(),
	UpdatedAt: time.Now(),
}

var ManyFarm = []model.Farm{
	SingleFarm,
	SingleFarm,
}
