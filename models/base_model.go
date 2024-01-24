package models

import "time"

type BaseModel struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	IsDeleted bool `gorm:"default:false"`
}
