package model

import (
	"time"
)

type Statistic struct {
	ID          uint      `json:"id" gorm:"primarykey"`
	RequestPath string    `json:"request_path"`
	Method      string    `json:"method"`
	UserAgent   string    `json:"user_agent"`
	Count       uint      `json:"count"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type GetMany struct {
	Page     int    `query:"page"`
	Quantity int    `query:"quantity"`
	Order    string `query:"order"`
}

type GetStatisticByUserAgent struct {
	UserAgent string `json:"user_agent" validate:"required" param:"id"`
}

// ================= RESPONSE =================
type StatisticResponse struct {
	ID          uint      `json:"id" gorm:"primarykey"`
	RequestPath string    `json:"request_path"`
	Method      string    `json:"method"`
	UserAgent   string    `json:"user_agent"`
	Count       uint      `json:"count"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
