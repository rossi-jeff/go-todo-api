package models

import "time"

type UserResponse struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	IsDeleted bool   `gorm:"default:false"`
	UserName  string `gorm:"size:25;unique;"`
	Email     string `gorm:"size:100"`
	Random    bool   `gorm:"default:false"`
}

type User struct {
	BaseModel
	UserName       string `gorm:"size:25;unique;"`
	Email          string `gorm:"size:100"`
	Random         bool   `gorm:"default:false"`
	PassWordDigest string `gorm:"size:255"`
	ToDos          []ToDo
}
