package models

import "time"

type Timestamps struct {
	CreatedAt time.Time  `gorm:"column:created_at;type:TIMESTAMP;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at;type:TIMESTAMP;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}