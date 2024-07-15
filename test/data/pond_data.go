package data

import (
	"time"

	model "github.com/harisspace/fisheries-api/modules/farm/models"
)

var SinglePond = model.Pond{
	ID:        1,
	FarmId:    "farmId",
	PondId:    "pondId",
	Name:      "Susi Farm",
	IsDeleted: false,
	CreatedAt: time.Now(),
	DeletedAt: time.Now(),
	UpdatedAt: time.Now(),
}

var ManyPond = []model.Pond{
	SinglePond,
	SinglePond,
}
