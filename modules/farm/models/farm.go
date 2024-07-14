package model

import "time"

type Farm struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement:true"`
	FarmId    string    `json:"farm_id" gorm:"unique"`
	Name      string    `json:"name"`
	IsDeleted bool      `json:"is_deleted"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
	Pond      []Pond    `json:"pond" gorm:"foreignKey:FarmId;references:FarmId"`
}

// Base
type GetMany struct {
	Page     int    `query:"page"`
	Quantity int    `query:"quantity"`
	Order    string `query:"order"`
}

// ========= COMMAND =========
type CreateFarm struct {
	Name string `json:"name" validate:"required"`
}

type GetManyFarm struct {
	GetMany
}

type GetManyPond struct {
	GetMany
}

type GetFarmById struct {
	FarmId string `json:"farm_id" validate:"required" param:"id"`
}

type UpdateFarm struct {
	FarmId string `json:"farm_id"`
	Name   string `json:"name"`
}

type DeleteFarm struct {
	FarmId string `json:"farm_id" param:"id"`
}

// ========= RESPONSE =========
type FarmResponse struct {
	FarmId    string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
