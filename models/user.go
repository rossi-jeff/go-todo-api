package models

type User struct {
	BaseModel
	UserName       string `gorm:"size:25;unique;"`
	PassWordDigest string `gorm:"size:255"`
	Email          string `gorm:"size:100"`
	Random         bool   `gorm:"default:false"`
	ToDos          []ToDo
}
