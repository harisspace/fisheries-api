package model

import "time"

type Pond struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement:true"`
	PondId    string    `json:"pond_id"`
	FarmId    string    `json:"farm_id"`
	Name      string    `json:"name"`
	IsDeleted bool      `json:"is_deleted"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

// ========= COMMAND =========
// TODO: handle type validation
type CreatePond struct {
	Name   string `json:"name" validate:"required"`
	FarmId string `json:"farm_id" validate:"required"`
}

type GetPondById struct {
	PondId string `json:"pond_id" validate:"required" param:"id"`
}

type UpdatePond struct {
	PondId string `json:"pond_id"`
	Name   string `json:"name"`
	FarmId string `json:"farm_id"`
}

type DeletePond struct {
	PondId string `json:"pond_id" param:"id"`
}

// ========= RESPONSE =========
type PondResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	FarmId    string    `json:"farm_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
